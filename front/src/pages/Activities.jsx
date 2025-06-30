import React, { useState, useEffect } from "react";
import "@styles/Activities.css";

const actividadesDataInicial = [
    {
        id: "1",
        nombre: "Surf",
        descripcion: "Clases para todos los niveles, desde principiantes hasta avanzados.",
        profesor: "Juan Pérez",
        horarios: "Lunes - Miércoles - Viernes 8:00 - 10:00",
        cupos: 15,
        foto: "/surf.png",
    },
    {
        id: "2",
        nombre: "Yoga",
        descripcion: "Yoga frente al mar para relajarte y mejorar tu flexibilidad.",
        profesor: "María López",
        horarios: "Martes - Jueves 18:00 - 19:30",
        cupos: 20,
        foto: "/yoga.png",
    },
    {
        id: "3",
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

    const [isEditing, setIsEditing] = useState(false);
    const [editandoActividad, setEditandoActividad] = useState(null);

    const [busqueda, setBusqueda] = useState("");
    const [inscripciones, setInscripciones] = useState(() => {
        const savedInscripciones = localStorage.getItem("userInscripciones");
        return savedInscripciones ? JSON.parse(savedInscripciones) : [];
    });

    // Guarda las inscripciones en localStorage cada vez que cambian
    useEffect(() => {
        localStorage.setItem("userInscripciones", JSON.stringify(inscripciones));
    }, [inscripciones]);


    useEffect(() => {
        if (actividadSeleccionada && !isEditing) {
            setEditandoActividad({ ...actividadSeleccionada });
        } else if (!actividadSeleccionada) {
            setEditandoActividad(null);
            setIsEditing(false); // Resetea el estado de edición si no hay actividad seleccionada
        }
    }, [actividadSeleccionada]);


    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const role = (localStorage.getItem("role") || "").toLowerCase(); // Normaliza el rol a minúsculas

    // Función para cerrar sesión
    const handleLogout = () => {
        localStorage.removeItem("isLoggedIn");
        localStorage.removeItem("role");
        localStorage.removeItem("userInscripciones");
        window.location.reload(); // Recarga la página para aplicar los cambios
    };

    const abrirDetalle = (actividad) => {
        setActividadSeleccionada(actividad);
        setEditandoActividad({ ...actividad }); // Inicializa el estado de edición con la actividad actual
        setShowForm(false);
        setIsEditing(false); // Abre el detalle en modo de vista por defecto
    };

    const cerrarDetalle = () => {
        setActividadSeleccionada(null);
        setEditandoActividad(null); // Limpia el estado de edición
        setShowForm(false);
        setIsEditing(false); // Asegura que el modo edición esté apagado
    };

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
        alert("La función de eliminar actividad no está disponible aún.");
    };

    // Maneja los cambios en los inputs del formulario de AGREGAR nueva actividad
    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setNuevaActividad((prev) => ({ ...prev, [name]: value }));
    };

    // Maneja los cambios en los inputs del formulario de EDICIÓN de actividad
    const handleEditInputChange = (e) => {
        const { name, value } = e.target;
        setEditandoActividad((prev) => ({ ...prev, [name]: value }));
    };

    const handleGuardarEdicion = () => {
        // Alerta de función no disponible para administrador
        alert("La función de guardar edición no está disponible aún.");
    };

    const handleCancelarEdicion = () => {
        // Alerta de función no disponible para administrador
        alert("La función de cancelar edición no está disponible aún.");
        setIsEditing(false); // Sale del modo edición
        setEditandoActividad({ ...actividadSeleccionada }); // Restaura la actividad original seleccionada
    };

    const handleAgregarActividad = (e) => {
        e.preventDefault(); // Previene el comportamiento por defecto del formulario
        // Alerta de función no disponible para administrador
        alert("La función de agregar actividad no está disponible aún.");
    };

    const handleBusquedaChange = (e) => {
        setBusqueda(e.target.value);
    };

    // Lógica de filtrado de actividades
    const actividadesFiltradasPorBusqueda = actividades.filter((act) => {
        const textoBusqueda = busqueda.toLowerCase();
        if (textoBusqueda.length === 0) return true; // Si la búsqueda está vacía, muestra todas

        // Filtra por las primeras 3 letras del nombre o del día en horarios
        const nombreSubstr = act.nombre.toLowerCase().substring(0, 3);
        const horariosSubstr = act.horarios.toLowerCase().substring(0, 3);

        return (
            nombreSubstr.startsWith(textoBusqueda) ||
            horariosSubstr.startsWith(textoBusqueda)
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
                    className="search-input"
                    aria-label="Buscar actividades por nombre o día"
                />

                {isLoggedIn && ( // Solo muestra "Cerrar Sesión" si isLoggedIn es true
                    <button className="btn-logout" onClick={handleLogout}>
                        Cerrar Sesión
                    </button>
                )}

                {isLoggedIn && role === "admin" && ( // Solo muestra "Agregar Actividad" si es admin
                    <button
                        className="btn-agregar"
                        onClick={() => { alert("La función de agregar actividad no está disponible aún."); setShowForm(false); }}
                    >
                        Agregar Nueva Actividad
                    </button>
                )}
            </div>

            {/* El formulario para agregar nueva actividad solo se muestra si showForm es true y es admin */}
            {showForm && isLoggedIn && role === "admin" && (
                <form
                    className="formulario-nueva-actividad"
                    onSubmit={handleAgregarActividad}
                >
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
                    <button type="submit" onClick={() => alert("La función de guardar actividad no está disponible aún.")}>
                        Guardar Actividad
                    </button>
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
                    <button
                        className="cerrar"
                        onClick={cerrarDetalle}
                        aria-label="Cerrar detalle"
                    >
                        &times;
                    </button>
                    {/* Campos de la actividad (editables o de solo lectura) */}
                    {isEditing && isLoggedIn && role === "admin" ? ( // Solo editable si es admin y está en modo edición
                        <input
                            type="text"
                            name="nombre"
                            value={editandoActividad.nombre}
                            onChange={handleEditInputChange}
                            className="formulario-edicion-input"
                        />
                    ) : (
                        <h2>{actividadSeleccionada.nombre}</h2>
                    )}

                    <div className="detalle-contenido">
                        {actividadSeleccionada.foto ? (
                            <img
                                src={isEditing && isLoggedIn && role === "admin" ? editandoActividad.foto : actividadSeleccionada.foto}
                                alt={isEditing && isLoggedIn && role === "admin" ? editandoActividad.nombre : actividadSeleccionada.nombre}
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
                                <strong>Descripción:</strong>{" "}
                                {isEditing && isLoggedIn && role === "admin" ? (
                                    <textarea
                                        name="descripcion"
                                        value={editandoActividad.descripcion}
                                        onChange={handleEditInputChange}
                                        className="formulario-edicion-input"
                                    />
                                ) : (
                                    actividadSeleccionada.descripcion
                                )}
                            </p>
                            <p>
                                <strong>Profesor:</strong>{" "}
                                {isEditing && isLoggedIn && role === "admin" ? (
                                    <input
                                        type="text"
                                        name="profesor"
                                        value={editandoActividad.profesor}
                                        onChange={handleEditInputChange}
                                        className="formulario-edicion-input"
                                    />
                                ) : (
                                    actividadSeleccionada.profesor
                                )}
                            </p>
                            <p>
                                <strong>Horarios:</strong>{" "}
                                {isEditing && isLoggedIn && role === "admin" ? (
                                    <input
                                        type="text"
                                        name="horarios"
                                        value={editandoActividad.horarios}
                                        onChange={handleEditInputChange}
                                        className="formulario-edicion-input"
                                    />
                                ) : (
                                    actividadSeleccionada.horarios
                                )}
                            </p>
                            <p>
                                <strong>Cupos:</strong>{" "}
                                {isEditing && isLoggedIn && role === "admin" ? (
                                    <input
                                        type="number"
                                        name="cupos"
                                        value={editandoActividad.cupos}
                                        onChange={handleEditInputChange}
                                        min={0}
                                        className="formulario-edicion-input"
                                    />
                                ) : (
                                    actividadSeleccionada.cupos
                                )}
                            </p>
                            {isEditing && isLoggedIn && role === "admin" && ( // Input de foto solo si es admin y está editando
                                <p>
                                    <strong>Foto URL:</strong>{" "}
                                    <input
                                        type="text"
                                        name="foto"
                                        value={editandoActividad.foto}
                                        onChange={handleEditInputChange}
                                        className="formulario-edicion-input"
                                    />
                                </p>
                            )}

                            {/* El botón de inscripción/dar de baja se muestra si isLoggedIn es true y NO estamos editando */}
                            {isLoggedIn && !isEditing && (
                                <button
                                    className="btn-inscribir"
                                    onClick={() => handleToggleInscription(actividadSeleccionada.id, actividadSeleccionada.nombre)}
                                >
                                    {inscripciones.includes(actividadSeleccionada.id)
                                        ? "Dar de Baja"
                                        : "Inscribirse"}
                                </button>
                            )}

                            {/* Mostrar botones de admin SOLO si isLoggedIn es true Y el role es "admin" */}
                            {isLoggedIn && role === "admin" && (
                                <div className="botones-admin">
                                    {isEditing ? (
                                        <>
                                            <button className="btn-editar" onClick={handleGuardarEdicion}>
                                                Guardar Cambios
                                            </button>
                                            <button className="btn-eliminar" onClick={handleCancelarEdicion}>
                                                Cancelar Edición
                                            </button>
                                        </>
                                    ) : (
                                        <>
                                            <button className="btn-editar" onClick={() => { setIsEditing(true); alert("La función de editar no está disponible aún."); }}>
                                                Editar Actividad
                                            </button>
                                            <button
                                                className="btn-eliminar"
                                                onClick={() => handleEliminarActividad(actividadSeleccionada.id)}
                                            >
                                                Eliminar Actividad
                                            </button>
                                        </>
                                    )}
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            )}

            {isLoggedIn && ( // Mostrar sección de "Mis Actividades Inscritas" solo si isLoggedIn es true
                <>
                    <hr className="divider" />
                    <h4 className="section-title">Mis Actividades Inscritas</h4>
                    <div className="mis-actividades-section">
                        {misActividadesInscritas.length > 0 ? (
                            <ul className="mis-inscripciones">
                                {misActividadesInscritas.map((act) => (
                                    <li key={act.id} style={{ position: 'relative' }}>
                                        {act.nombre}
                                        {isLoggedIn && ( // El botón de dar de baja inline también requiere isLoggedIn
                                            <button
                                                className="btn-dar-baja-inline"
                                                onClick={() => handleToggleInscription(act.id, act.nombre)}
                                                aria-label={`Dar de baja de ${act.nombre}`}
                                            >
                                                &times;
                                            </button>
                                        )}
                                    </li>
                                ))}
                            </ul>
                        ) : (
                            <p className="no-inscripciones-msg">No estás inscripto en ninguna actividad.</p>
                        )}
                    </div>
                </>
            )}
        </section>
    );
};

export default Activities;