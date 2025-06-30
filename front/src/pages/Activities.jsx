import React, { useState, useEffect } from "react";
import "@styles/Activities.css";

const Activities = () => {
    const [actividades, setActividades] = useState([]);
    const [actividadSeleccionada, setActividadSeleccionada] = useState(null);
    const [showForm, setShowForm] = useState(false);
    const [nuevaActividad, setNuevaActividad] = useState({
        title: "",
        description: "",
        day: "",
        duration: "",
        category: "",
        state: "",
        instructor: "",
        cupo: "",
    });

    const [isEditing, setIsEditing] = useState(false);
    const [editandoActividad, setEditandoActividad] = useState(null);

    const [busqueda, setBusqueda] = useState("");
    const [inscripciones, setInscripciones] = useState(() => {
        const savedInscripciones = localStorage.getItem("userInscripciones");
        return savedInscripciones ? JSON.parse(savedInscripciones) : [];
    });

    useEffect(() => {
        fetch("http://localhost:8080/activities")
            .then((res) => res.json())
            .then((data) => setActividades(data))
            .catch((err) => {
                console.error("Error al obtener actividades:", err);
                alert("No se pudieron cargar las actividades.");
            });
    }, []);

    useEffect(() => {
        localStorage.setItem("userInscripciones", JSON.stringify(inscripciones));
    }, [inscripciones]);

    useEffect(() => {
        if (actividadSeleccionada && !isEditing) {
            setEditandoActividad({ ...actividadSeleccionada });
        } else if (!actividadSeleccionada) {
            setEditandoActividad(null);
            setIsEditing(false);
        }
    }, [actividadSeleccionada]);

    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const role = (localStorage.getItem("role") || "").toLowerCase();

    const handleLogout = () => {
        localStorage.clear();
        window.location.reload();
    };

    const abrirDetalle = (actividad) => {
        setActividadSeleccionada(actividad);
        setEditandoActividad({ ...actividad });
        setShowForm(false);
        setIsEditing(false);
    };

    const cerrarDetalle = () => {
        setActividadSeleccionada(null);
        setEditandoActividad(null);
        setShowForm(false);
        setIsEditing(false);
    };

    const handleToggleInscription = (actividadId, nombreActividad) => {
        if (inscripciones.includes(actividadId)) {
            if (window.confirm(`¿Dar de baja de ${nombreActividad}?`)) {
                setInscripciones(inscripciones.filter((id) => id !== actividadId));
                alert(`Baja realizada: ${nombreActividad}`);
            }
        } else {
            setInscripciones([...inscripciones, actividadId]);
            alert(`Inscripción exitosa: ${nombreActividad}`);
        }
    };

    const handleBusquedaChange = (e) => setBusqueda(e.target.value);

    const actividadesFiltradasPorBusqueda = actividades.filter((act) => {
        const texto = busqueda.toLowerCase();
        if (!texto) return true;

        return (
            act.title.toLowerCase().startsWith(texto) ||
            act.day.toLowerCase().startsWith(texto) ||
            act.category.toLowerCase().startsWith(texto)
        );
    });


    const misActividadesInscritas = actividades.filter((act) =>
        inscripciones.includes(act.ID?.toString())
    );
    const handleAgregarActividad = async (e) => {
        e.preventDefault();

        // Validación mínima
        if (!nuevaActividad.title || !nuevaActividad.day || !nuevaActividad.category) {
            alert("Por favor completá los campos obligatorios.");
            return;
        }

        const actividadFormateada = {
            title: nuevaActividad.title,
            description: nuevaActividad.description,
            day: nuevaActividad.day,
            duration: parseInt(nuevaActividad.duration, 10),
            category: nuevaActividad.category,
            state: nuevaActividad.state || "Activo",
            instructor: nuevaActividad.instructor,
            cupo: parseInt(nuevaActividad.cupo, 10),
        };

        try {
            const response = await fetch("http://localhost:8080/admin/activity", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(actividadFormateada),
            });

            if (!response.ok) {
                throw new Error("Error al crear la actividad");
            }

            const nueva = await response.json();

            setActividades((prev) => [...prev, nueva]);
            setShowForm(false);
            setNuevaActividad({
                title: "",
                description: "",
                day: "",
                duration: "",
                category: "",
                state: "",
                instructor: "",
                cupo: "",
            });

            alert("Actividad agregada con éxito.");
        } catch (error) {
            console.error("Error:", error);
            alert("No se pudo agregar la actividad.");
        }
    };



    return (
        <section className="actividades">
            <h3>Nuestras Actividades</h3>

            <div className="controles-actividades">
                <input
                    type="text"
                    placeholder="Buscar actividad por nombre o día..."
                    value={busqueda}
                    onChange={handleBusquedaChange}
                    className="search-input"
                />

                {isLoggedIn && (
                    <button className="btn-logout" onClick={handleLogout}>
                        Cerrar Sesión
                    </button>
                )}

                {isLoggedIn && role === "admin" && (
                    <button
                        className="btn-agregar"
                        onClick={() => {
                            setShowForm(true);
                        }}
                    >
                        Agregar Nueva Actividad
                    </button>
                )}
            </div>
            {showForm && isLoggedIn && role === "admin" && (
                <form className="formulario-nueva-actividad" onSubmit={handleAgregarActividad}>
                    <input
                        type="text"
                        name="title"
                        placeholder="Título"
                        value={nuevaActividad.title}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, title: e.target.value })}
                        required
                    />
                    <textarea
                        name="description"
                        placeholder="Descripción"
                        value={nuevaActividad.description}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, description: e.target.value })}
                    />
                    <input
                        type="text"
                        name="day"
                        placeholder="Día"
                        value={nuevaActividad.day}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, day: e.target.value })}
                        required
                    />
                    <input
                        type="number"
                        name="duration"
                        placeholder="Duración (minutos)"
                        value={nuevaActividad.duration}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, duration: e.target.value })}
                        required
                    />
                    <input
                        type="text"
                        name="category"
                        placeholder="Categoría"
                        value={nuevaActividad.category}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, category: e.target.value })}
                        required
                    />
                    <input
                        type="text"
                        name="state"
                        placeholder="Estado"
                        value={nuevaActividad.state}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, state: e.target.value })}
                    />
                    <input
                        type="text"
                        name="instructor"
                        placeholder="Instructor"
                        value={nuevaActividad.instructor}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, instructor: e.target.value })}
                    />
                    <input
                        type="number"
                        name="cupo"
                        placeholder="Cupo"
                        value={nuevaActividad.cupo}
                        onChange={(e) => setNuevaActividad({ ...nuevaActividad, cupo: e.target.value })}
                        required
                    />
                    <button type="submit">Guardar Actividad</button>
                </form>
            )}

            <div className="actividades-grid">
                {actividadesFiltradasPorBusqueda.length > 0 ? (
                    actividadesFiltradasPorBusqueda.map((act) => (
                        <div
                            key={act.ID}
                            className="actividad"
                            onClick={() => abrirDetalle(act)}
                            tabIndex={0}
                            role="button"
                            onKeyDown={(e) => e.key === "Enter" && abrirDetalle(act)}
                        >
                            {act.title}
                        </div>
                    ))
                ) : (
                    <p>No se encontraron actividades.</p>
                )}
            </div>

            {actividadSeleccionada && (
                <div className="detalle-actividad">
                    <button className="cerrar" onClick={cerrarDetalle}>
                        &times;
                    </button>
                    <h2>{actividadSeleccionada.title}</h2>

                    <div className="detalle-contenido">
                        <div className="detalle-texto">
                            <p><strong>Descripción:</strong> {actividadSeleccionada.description}</p>
                            <p><strong>Día:</strong> {actividadSeleccionada.day}</p>
                            <p><strong>Duración:</strong> {actividadSeleccionada.duration} min</p>
                            <p><strong>Categoría:</strong> {actividadSeleccionada.category}</p>
                            <p><strong>Estado:</strong> {actividadSeleccionada.state}</p>
                            <p><strong>Profesor:</strong> {actividadSeleccionada.instructor}</p>
                            <p><strong>Cupo:</strong> {actividadSeleccionada.cupo}</p>

                            {isLoggedIn && !isEditing && (
                                <button
                                    className="btn-inscribir"
                                    onClick={() =>
                                        handleToggleInscription(
                                            actividadSeleccionada.ID?.toString(),
                                            actividadSeleccionada.title
                                        )
                                    }
                                >
                                    {inscripciones.includes(actividadSeleccionada.ID?.toString())
                                        ? "Dar de Baja"
                                        : "Inscribirse"}
                                </button>
                            )}
                        </div>
                    </div>
                </div>
            )}

            {isLoggedIn && (
                <>
                    <hr className="divider" />
                    <h4 className="section-title">Mis Actividades Inscritas</h4>
                    <div className="mis-actividades-section">
                        {misActividadesInscritas.length > 0 ? (
                            <ul className="mis-inscripciones">
                                {misActividadesInscritas.map((act) => (
                                    <li key={act.ID}>
                                        {act.title}
                                        <button
                                            className="btn-dar-baja-inline"
                                            onClick={() =>
                                                handleToggleInscription(
                                                    act.ID?.toString(),
                                                    act.title
                                                )
                                            }
                                        >
                                            &times;
                                        </button>
                                    </li>
                                ))}
                            </ul>
                        ) : (
                            <p className="no-inscripciones-msg">
                                No estás inscripto en ninguna actividad.
                            </p>
                        )}
                    </div>
                </>
            )}
        </section>
    );
};

export default Activities;
