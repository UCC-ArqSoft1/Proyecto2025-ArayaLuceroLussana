import React, { useState, useEffect } from "react";
import "./Activities.css";

const actividadesDataInicial = [
    {
        id: "surf",
        nombre: "Surf",
        descripcion: "Clases para todos los niveles, desde principiantes hasta avanzados.",
        profesor: "Juan Pérez",
        horarios: "Lunes, Miércoles y Viernes 8:00 - 10:00",
        cupos: 15,
        foto: "/surf.png",
    },
    {
        id: "yoga",
        nombre: "Yoga",
        descripcion: "Yoga frente al mar para relajarte y mejorar tu flexibilidad.",
        profesor: "María López",
        horarios: "Martes y Jueves 18:00 - 19:30",
        cupos: 20,
        foto: "/yoga.png",
    },
    {
        id: "skate",
        nombre: "Skate",
        descripcion: "Skate para todos, con técnicas y trucos para principiantes y expertos.",
        profesor: "Carlos Gómez",
        horarios: "Sábados 10:00 - 12:00",
        cupos: 12,
        foto: "/skate.png",
    },
];

const Activities = () => {
    const [actividades, setActividades] = useState(actividadesDataInicial);
    const [actividadSeleccionada, setActividadSeleccionada] = useState(null);
    const [showForm, setShowForm] = useState(false);
    const [nuevaActividad, setNuevaActividad] = useState({
        id: "",
        nombre: "",
        descripcion: "",
        profesor: "",
        horarios: "",
        cupos: "",
        foto: "",
    });

    const [busqueda, setBusqueda] = useState("");
    const [inscripciones, setInscripciones] = useState(() => {
        const savedInscripciones = localStorage.getItem("userInscripciones");
        return savedInscripciones ? JSON.parse(savedInscripciones) : [];
    });

    // Efecto para guardar las inscripciones en localStorage cada vez que cambian
    useEffect(() => {
        localStorage.setItem("userInscripciones", JSON.stringify(inscripciones));
    }, [inscripciones]);

    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const role = localStorage.getItem("role") || "";

    const abrirDetalle = (actividad) => {
        setActividadSeleccionada(actividad);
        setShowForm(false);
    };

    const cerrarDetalle = () => {
        setActividadSeleccionada(null);
        setShowForm(false);
    };

    // Función para manejar inscripción y desinscripción
    const handleToggleInscription = (actividadId, nombreActividad) => {
        if (inscripciones.includes(actividadId)) {
            if (window.confirm(`¿Estás seguro de que quieres dar de baja de ${nombreActividad}?`)) {
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
            setInscripciones(inscripciones.filter((inscId) => inscId !== id)); // Eliminar de inscripciones también
            cerrarDetalle();
        }
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setNuevaActividad((prev) => ({ ...prev, [name]: value }));
    };

    const handleAgregarActividad = (e) => {
        e.preventDefault();
        if (!nuevaActividad.id.trim() || !nuevaActividad.nombre.trim()) {
            alert("El ID y nombre son obligatorios.");
            return;
        }
        if (actividades.some((act) => act.id === nuevaActividad.id.trim())) {
            alert("Ya existe una actividad con ese ID.");
            return;
        }
        setActividades([
            ...actividades,
            { ...nuevaActividad, cupos: Number(nuevaActividad.cupos) || 0 },
        ]);
        setNuevaActividad({
            id: "",
            nombre: "",
            descripcion: "",
            profesor: "",
            horarios: "",
            cupos: "",
            foto: "",
        });
        setShowForm(false);
    };

    const handleBusquedaChange = (e) => {
        setBusqueda(e.target.value);
    };

    const actividadesFiltradasPorBusqueda = actividades.filter((act) => {
        const textoBusqueda = busqueda.toLowerCase();
        return (
            act.nombre.toLowerCase().includes(textoBusqueda) ||
            act.horarios.toLowerCase().includes(textoBusqueda)
        );
    });

    // Actividades a las que el usuario está inscrito
    const misActividadesInscritas = actividades.filter((act) =>
        inscripciones.includes(act.id)
    );

    return (
        <section className="actividades">
            <h3>Nuestras Actividades</h3>

            {/* Controles: Búsqueda y botones de acción */}
            <div className="controles-actividades">
                <input
                    type="text"
                    placeholder="Buscar por nombre o día..."
                    value={busqueda}
                    onChange={handleBusquedaChange}
                    className="input-busqueda"
                    aria-label="Buscar actividades por nombre o día"
                />

                {isLoggedIn && role === "admin" && (
                    <button className="btn-agregar" onClick={() => setShowForm((prev) => !prev)}>
                        {showForm ? "Cancelar" : "Agregar Nueva Actividad"}
                    </button>
                )}
            </div>

            {/* Formulario para agregar nueva actividad (solo admin) */}
            {showForm && (
                <form className="formulario-nueva-actividad" onSubmit={handleAgregarActividad}>
                    <input
                        type="text"
                        name="id"
                        placeholder="ID única (ej: surf)"
                        value={nuevaActividad.id}
                        onChange={handleInputChange}
                        required
                    />
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
                        name="horarios"
                        placeholder="Días / Horarios"
                        value={nuevaActividad.horarios}
                        onChange={handleInputChange}
                    />
                    <input
                        type="number"
                        name="cupos"
                        placeholder="Cupos disponibles"
                        value={nuevaActividad.cupos}
                        onChange={handleInputChange}
                        min={0}
                    />
                    <input
                        type="text"
                        name="foto"
                        placeholder="URL de la foto"
                        value={nuevaActividad.foto}
                        onChange={handleInputChange}
                    />
                    <button type="submit">Guardar Actividad</button>
                </form>
            )}

            {/* Grilla de todas las actividades (filtradas por búsqueda) */}
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

            {/* --- Detalle de la actividad seleccionada y Mis Actividades --- */}
            {actividadSeleccionada && (
                <div className="detalle-actividad">
                    <button className="cerrar" onClick={cerrarDetalle} aria-label="Cerrar detalle">
                        &times;
                    </button>
                    <h2>{actividadSeleccionada.nombre}</h2>
                    <div className="detalle-contenido">
                        {actividadSeleccionada.foto ? (
                            <img
                                src={actividadSeleccionada.foto}
                                alt={actividadSeleccionada.nombre}
                                className="foto-actividad"
                            />
                        ) : (
                            <div
                                className="foto-actividad"
                                style={{
                                    backgroundColor: "#ccc",
                                    width: "200px",
                                    height: "120px",
                                    borderRadius: "6px",
                                    display: "flex",
                                    alignItems: "center",
                                    justifyContent: "center",
                                    color: "#666",
                                }}
                            >
                                Sin imagen
                            </div>
                        )}
                        <div className="detalle-texto">
                            <p>
                                <strong>Descripción:</strong> {actividadSeleccionada.descripcion}
                            </p>
                            <p>
                                <strong>Profesor:</strong> {actividadSeleccionada.profesor}
                            </p>
                            <p>
                                <strong>Días / Horarios:</strong> {actividadSeleccionada.horarios}
                            </p>
                            <p>
                                <strong>Cupos disponibles:</strong> {actividadSeleccionada.cupos}
                            </p>

                            {isLoggedIn && role === "user" && (
                                <button
                                    className={inscripciones.includes(actividadSeleccionada.id) ? "btn-dar-baja" : "btn-inscribir"}
                                    onClick={() =>
                                        handleToggleInscription(
                                            actividadSeleccionada.id,
                                            actividadSeleccionada.nombre
                                        )
                                    }
                                >
                                    {inscripciones.includes(actividadSeleccionada.id) ? "Dar de Baja" : "Inscribirse"}
                                </button>
                            )}

                            {isLoggedIn && role === "admin" && (
                                <div className="botones-admin">
                                    <button
                                        className="btn-eliminar"
                                        onClick={() => handleEliminarActividad(actividadSeleccionada.id)}
                                    >
                                        Eliminar actividad
                                    </button>
                                    <button
                                        className="btn-editar"
                                        onClick={() => alert("Función de editar no implementada aún")}
                                    >
                                        Editar actividad
                                    </button>
                                </div>
                            )}
                        </div>
                    </div>

                    {/* --- Sección de Mis Actividades ahora dentro del detalle --- */}
                    {isLoggedIn && role === "user" && (
                        <>
                            <hr className="divider" /> {/* Divisor visual */}
                            <h3 className="section-title">Mis Actividades Inscritas</h3>
                            {misActividadesInscritas.length > 0 ? (
                                <div className="mis-actividades-grid">
                                    {misActividadesInscritas.map((act) => (
                                        <div
                                            key={`my-${act.id}`} // Usar una key única para evitar conflictos
                                            className="actividad my-actividad"
                                            onClick={() => abrirDetalle(act)}
                                            tabIndex={0}
                                            role="button"
                                            onKeyDown={(e) => {
                                                if (e.key === "Enter") abrirDetalle(act);
                                            }}
                                        >
                                            {act.nombre}
                                            <button
                                                className="btn-dar-baja-inline"
                                                onClick={(e) => {
                                                    e.stopPropagation(); // Evita abrir el detalle al dar de baja
                                                    handleToggleInscription(act.id, act.nombre);
                                                }}
                                                aria-label={`Dar de baja de ${act.nombre}`}
                                            >
                                                &times;
                                            </button>
                                        </div>
                                    ))}
                                </div>
                            ) : (
                                <p className="no-inscripciones-msg">No estás inscrito en ninguna actividad aún.</p>
                            )}
                        </>
                    )}
                    {/* --- Fin Sección de Mis Actividades (dentro del detalle) --- */}
                </div>
            )}
        </section>
    );
};

export default Activities;