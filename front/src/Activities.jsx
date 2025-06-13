import React, { useState, useEffect } from "react";
import "./Activities.css";

const Activities = () => {
    const [actividades, setActividades] = useState([
        {
            id: "1",
            nombre: "Yoga",
            descripcion: "Clase de Yoga relajante",
            profesor: "Ana",
            dia: "Lunes",
            hora: "10:00",
            duration: 60,
            category: "Fitness",
            cupos: 15,
            state: "Activo",
        },
    ]);
    const [actividadSeleccionada, setActividadSeleccionada] = useState(null);
    const [showForm, setShowForm] = useState(false);
    const [nuevaActividad, setNuevaActividad] = useState({
        nombre: "",
        descripcion: "",
        profesor: "",
        dia: "",
        hora: "",
        duration: "",
        category: "",
        cupos: ""
    });

    const [busqueda, setBusqueda] = useState("");
    const [inscripciones, setInscripciones] = useState(() => {
        const savedInscripciones = localStorage.getItem("userInscripciones");
        return savedInscripciones ? JSON.parse(savedInscripciones) : [];
    });

    useEffect(() => {
        localStorage.setItem("userInscripciones", JSON.stringify(inscripciones));
    }, [inscripciones]);

    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const role = localStorage.getItem("role") || "";

    const handleLogout = () => {
        localStorage.removeItem("isLoggedIn");
        localStorage.removeItem("role");
        localStorage.removeItem("userInscripciones");
        window.location.reload();
    };

    const abrirDetalle = (actividad) => {
        setActividadSeleccionada(actividad);
        setShowForm(false);
    };

    const cerrarDetalle = () => {
        setActividadSeleccionada(null);
        setShowForm(false);
    };

    const handleToggleInscription = (actividadId, nombreActividad) => {
        if (inscripciones.includes(actividadId)) {
            if (window.confirm(`¿Estás seguro de que deseas dar de baja de ${nombreActividad}?`)) {
                setInscripciones(inscripciones.filter((id) => id !== actividadId));
                alert(`Has dado de baja de: ${nombreActividad}`);
            }
        } else {
            setInscripciones([...inscripciones, actividadId]);
            alert(`Te has inscrito en: ${nombreActividad}`);
        }
    };

    const handleEliminarActividad = (id) => {
        if (window.confirm("¿Seguro que deseas eliminar esta actividad?")) {
            setActividades(actividades.filter((act) => act.id !== id));
            setInscripciones(inscripciones.filter((inscId) => inscId !== id));
            cerrarDetalle();
        }
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setNuevaActividad((prev) => ({
            ...prev,
            [name]: value,
        }));
    };

    const handleAgregarActividad = async (e) => {
        e.preventDefault();

        if (!nuevaActividad.nombre.trim()) {
            alert("El nombre de la actividad es obligatorio.");
            return;
        }
        if (!nuevaActividad.dia.trim()) {
            alert("El día es obligatorio.");
            return;
        }
        if (!nuevaActividad.hora.trim()) {
            alert("La hora es obligatoria.");
            return;
        }

        const token = localStorage.getItem("token");
        if (!token) {
            alert("No estás autenticado. Por favor inicia sesión.");
            return;
        }

        try {
            const activityToSave = {
                title: nuevaActividad.nombre,
                description: nuevaActividad.descripcion,
                day: nuevaActividad.dia,
                time: nuevaActividad.hora,
                duration: Number(nuevaActividad.duration) || 0,
                category: nuevaActividad.category || "General",
                state: "Activo",
                instructor: nuevaActividad.profesor,
                cupo: Number(nuevaActividad.cupos) || 0,
            };

            const response = await fetch("http://localhost:8080/admin/activity", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`
                },
                body: JSON.stringify(activityToSave),
            });

            if (response.ok) {
                alert("Actividad creada con éxito.");

                setActividades((prev) => [
                    ...prev,
                    {
                        id: crypto.randomUUID(),
                        nombre: nuevaActividad.nombre,
                        descripcion: nuevaActividad.descripcion,
                        profesor: nuevaActividad.profesor,
                        dia: nuevaActividad.dia,
                        hora: nuevaActividad.hora,
                        duration: Number(nuevaActividad.duration),
                        category: nuevaActividad.category || "General",
                        cupos: Number(nuevaActividad.cupos),
                        state: "Activo",
                    },
                ]);

                setNuevaActividad({
                    nombre: "",
                    descripcion: "",
                    profesor: "",
                    dia: "",
                    hora: "",
                    duration: "",
                    category: "",
                    cupos: ""
                });

                setShowForm(false);
            } else {
                const error = await response.json();
                console.error(error);
                alert("Error al crear la actividad.");
            }
        } catch (error) {
            console.error(error);
            alert("Error al crear la actividad.");
        }
    };


    const handleBusquedaChange = (e) => {
        setBusqueda(e.target.value);
    };

    const actividadesFiltradasPorBusqueda = actividades.filter((act) => {
        const textoBusqueda = busqueda.toLowerCase();
        if (textoBusqueda.length === 0) return true;
        return (
            act.nombre.toLowerCase().startsWith(textoBusqueda) ||
            act.dia.toLowerCase().startsWith(textoBusqueda)
        );
    });

    const misActividadesInscritas = actividades.filter((act) =>
        inscripciones.includes(act.id)
    );

    return (
        <section className="actividades">
            <h3>Nuestras Actividades</h3>

            <div className="controles-actividades">
                <input
                    type="text"
                    placeholder="Buscar actividad por nombre o día..."
                    value={busqueda}
                    onChange={handleBusquedaChange}
                    className="input-busqueda"
                    aria-label="Buscar actividades por nombre o día"
                />

                {isLoggedIn && (
                    <button className="btn-logout" onClick={handleLogout}>
                        Cerrar Sesión
                    </button>
                )}

                {isLoggedIn && role === "Admin" && (
                    <button
                        className="btn-agregar"
                        onClick={() => setShowForm((prev) => !prev)}
                    >
                        {showForm ? "Cancelar" : "Agregar Nueva Actividad"}
                    </button>
                )}
            </div>

            {showForm && (
                <form className="formulario-nueva-actividad" onSubmit={handleAgregarActividad}>
                    <input
                        type="text"
                        name="nombre"
                        placeholder="Nombre de la actividad"
                        value={nuevaActividad.nombre}
                        onChange={handleInputChange}
                        required
                    />
                    <textarea
                        name="descripcion"
                        placeholder="Descripción"
                        value={nuevaActividad.descripcion}
                        onChange={handleInputChange}
                    ></textarea>
                    <input
                        type="text"
                        name="profesor"
                        placeholder="Profesor"
                        value={nuevaActividad.profesor}
                        onChange={handleInputChange}
                    />
                    <input
                        type="text"
                        name="dia"
                        placeholder="Día (ej: Lunes)"
                        value={nuevaActividad.dia}
                        onChange={handleInputChange}
                        required
                    />
                    <input
                        type="time"
                        name="hora"
                        placeholder="Hora"
                        value={nuevaActividad.hora}
                        onChange={handleInputChange}
                        required
                    />
                    <input
                        type="number"
                        name="duration"
                        placeholder="Duración (minutos)"
                        value={nuevaActividad.duration}
                        onChange={handleInputChange}
                        min={1}
                    />
                    <input
                        type="number"
                        name="cupos"
                        placeholder="Cupos disponibles"
                        value={nuevaActividad.cupos}
                        onChange={handleInputChange}
                        min={1}
                    />
                    <input
                        type="text"
                        name="category"
                        placeholder="Categoría (ej: Fitness, Aire Libre)"
                        value={nuevaActividad.category}
                        onChange={handleInputChange}
                    />
                    <button type="submit">Guardar Actividad</button>
                </form>
            )}

            <div className="actividades-grid">
                {actividadesFiltradasPorBusqueda.length > 0 ? (
                    actividadesFiltradasPorBusqueda.map((act) => (
                        <div
                            key={act.id}
                            className="actividad"
                            onClick={() => abrirDetalle(act)}
                            tabIndex={0}
                            role="button"
                            onKeyDown={(e) => {
                                if (e.key === "Enter") abrirDetalle(act);
                            }}
                        >
                            {act.nombre}
                        </div>
                    ))
                ) : (
                    <p>No se encontraron actividades que coincidan con la búsqueda.</p>
                )}
            </div>

            {actividadSeleccionada && (
                <div className="detalle-actividad">
                    <button className="cerrar" onClick={cerrarDetalle} aria-label="Cerrar detalle">
                        &times;
                    </button>
                    <h2>{actividadSeleccionada.nombre}</h2>
                    <div className="detalle-contenido">
                        <div className="detalle-texto">
                            <p><strong>Descripción:</strong> {actividadSeleccionada.descripcion}</p>
                            <p><strong>Profesor:</strong> {actividadSeleccionada.profesor}</p>
                            <p><strong>Horario:</strong> {actividadSeleccionada.dia} {actividadSeleccionada.hora}</p>
                            <p><strong>Duración:</strong> {actividadSeleccionada.duration} minutos</p>
                            <p><strong>Cupos:</strong> {actividadSeleccionada.cupos}</p>
                            <p><strong>Categoría:</strong> {actividadSeleccionada.category}</p>
                            <p><strong>Estado:</strong> {actividadSeleccionada.state}</p>

                            {isLoggedIn && (
                                <button
                                    className="btn-inscripcion"
                                    onClick={() =>
                                        handleToggleInscription(
                                            actividadSeleccionada.id,
                                            actividadSeleccionada.nombre
                                        )
                                    }
                                >
                                    {inscripciones.includes(actividadSeleccionada.id)
                                        ? "Dar de Baja"
                                        : "Inscribirse"}
                                </button>
                            )}

                            {isLoggedIn && role === "Admin" && (
                                <button
                                    className="btn-eliminar"
                                    onClick={() => handleEliminarActividad(actividadSeleccionada.id)}
                                >
                                    Eliminar Actividad
                                </button>
                            )}
                        </div>
                    </div>
                </div>
            )}

            {isLoggedIn && (
                <>
                    <h4>Mis Actividades Inscritas</h4>
                    {misActividadesInscritas.length > 0 ? (
                        <ul className="mis-inscripciones">
                            {misActividadesInscritas.map((act) => (
                                <li key={act.id}>{act.nombre}</li>
                            ))}
                        </ul>
                    ) : (
                        <p>No estás inscripto en ninguna actividad.</p>
                    )}
                </>
            )}
        </section>
    );
};

export default Activities;
