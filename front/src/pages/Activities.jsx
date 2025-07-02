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
        cupo: ""
    });

    const [isEditing, setIsEditing] = useState(false);
    const [editandoActividad, setEditandoActividad] = useState(null);

    const [busqueda, setBusqueda] = useState("");
    const [inscripciones, setInscripciones] = useState([{}]);

    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    const role = (localStorage.getItem("role") || "").toLowerCase();
    const userId = localStorage.getItem("userId");

    // Solo carga todas las actividades, sin filtrar por usuario
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
        if (actividadSeleccionada && !isEditing) {
            setEditandoActividad({ ...actividadSeleccionada });
        } else if (!actividadSeleccionada) {
            setEditandoActividad(null);
            setIsEditing(false);
        }
    }, [actividadSeleccionada]);

    const handleToggleInscription = async (actividadId, nombreActividad) => {
        const userId = localStorage.getItem("userId");

        if (!userId) {
            alert("Usuario no identificado. Por favor inicia sesión.");
            return;
        }

        try {
            if (inscripciones.includes(actividadId)) {
                if (window.confirm(`¿Dar de baja de ${nombreActividad}?`)) {
                    const res = await fetch(`http://localhost:8080/socio/inscription/${actividadId}/${userId}`, {
                        method: "DELETE",
                        headers: {
                            "Content-Type": "application/json",
                            Authorization: `Bearer ${localStorage.getItem("token")}`,
                        },
                    });
                    if (!res.ok) throw new Error("Error al dar de baja");

                    setInscripciones(inscripciones.filter((id) => id !== actividadId));
                    alert(`Baja realizada: ${nombreActividad}`);
                }
            } else {
                const res = await fetch(`http://localhost:8080/socio/enroll/${userId}/${actividadId}`, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${localStorage.getItem("token")}`,
                    },
                    body: JSON.stringify({ actividadId, userId }),
                });

                if (!res.ok) {
                    const errorData = await res.json();
                    throw new Error(errorData.message || "Error al inscribir");
                }

                setInscripciones([...inscripciones, actividadId]);
                alert(`Inscripción exitosa: ${nombreActividad}`);
            }
        } catch (error) {
            console.error(error);
            alert(error.message);
        }
    };

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
                    Role: localStorage.getItem("role") || "Admin",
                },
                body: JSON.stringify(actividadFormateada),
            });

            if (!response.ok) {
                const errorData = await response.json();
                const msg = errorData.message || "Error al crear la actividad";
                throw new Error(msg);
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
            alert("No se pudo agregar la actividad: " + error.message);
        }
    };

    const handleEditarActividad = async () => {
        if (!editandoActividad || !editandoActividad.ID) {
            alert("Actividad inválida para editar.");
            return;
        }

        try {
            const response = await fetch(`http://localhost:8080/admin/activity/${editandoActividad.ID}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Role: localStorage.getItem("role") || "Admin",
                },
                body: JSON.stringify(editandoActividad),
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(result.message || "Error al editar la actividad");
            }

            setActividades((prev) =>
                prev.map((act) => (act.ID === editandoActividad.ID ? { ...editandoActividad } : act))
            );
            setActividadSeleccionada({ ...editandoActividad });
            setIsEditing(false);
            alert("Actividad actualizada exitosamente.");
        } catch (error) {
            console.error("Error al editar:", error);
            alert(error.message);
        }
    };

    const handleEliminarActividad = async (id) => {
        if (!window.confirm("¿Confirmás que querés eliminar esta actividad?")) return;

        try {
            const response = await fetch(`http://localhost:8080/admin/activity/${id}`, {
                method: "DELETE",
                headers: {
                    Role: localStorage.getItem("role") || "Admin",
                },
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(result.message || "Error al eliminar");
            }

            setActividades((prev) => prev.filter((act) => act.ID !== id));
            setActividadSeleccionada(null);
            alert("Actividad eliminada exitosamente.");
        } catch (error) {
            console.error("Error al eliminar:", error);
            alert(error.message);
        }
    }

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
                    <h2>{isEditing ? (
                        <input
                            type="text"
                            value={editandoActividad.title}
                            onChange={(e) =>
                                setEditandoActividad({ ...editandoActividad, title: e.target.value })
                            }
                        />
                    ) : (
                        actividadSeleccionada.title
                    )}</h2>

                    <div className="detalle-contenido">
                        <div className="detalle-texto">
                            <p>
                                <strong>Descripción:</strong>{" "}
                                {isEditing ? (
                                    <textarea
                                        value={editandoActividad.description}
                                        onChange={(e) =>
                                            setEditandoActividad({ ...editandoActividad, description: e.target.value })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.description
                                )}
                            </p>
                            <p>
                                <strong>Día:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="text"
                                        value={editandoActividad.day}
                                        onChange={(e) =>
                                            setEditandoActividad({ ...editandoActividad, day: e.target.value })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.day
                                )}
                            </p>
                            <p>
                                <strong>Duración:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="number"
                                        value={editandoActividad.duration}
                                        onChange={(e) =>
                                            setEditandoActividad({
                                                ...editandoActividad,
                                                duration: Number(e.target.value),
                                            })
                                        }
                                    />
                                ) : (
                                    `${actividadSeleccionada.duration} min`
                                )}
                            </p>
                            <p>
                                <strong>Categoría:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="text"
                                        value={editandoActividad.category}
                                        onChange={(e) =>
                                            setEditandoActividad({ ...editandoActividad, category: e.target.value })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.category
                                )}
                            </p>
                            <p>
                                <strong>Estado:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="text"
                                        value={editandoActividad.state}
                                        onChange={(e) =>
                                            setEditandoActividad({ ...editandoActividad, state: e.target.value })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.state
                                )}
                            </p>
                            <p>
                                <strong>Profesor:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="text"
                                        value={editandoActividad.instructor}
                                        onChange={(e) =>
                                            setEditandoActividad({ ...editandoActividad, instructor: e.target.value })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.instructor
                                )}
                            </p>
                            <p>
                                <strong>Cupo:</strong>{" "}
                                {isEditing ? (
                                    <input
                                        type="number"
                                        value={editandoActividad.cupo}
                                        onChange={(e) =>
                                            setEditandoActividad({
                                                ...editandoActividad,
                                                cupo: Number(e.target.value),
                                            })
                                        }
                                    />
                                ) : (
                                    actividadSeleccionada.cupo
                                )}
                            </p>

                            {!isEditing && isLoggedIn && (
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

                        {isLoggedIn && role === "admin" && (
                            <div className="botones-admin">
                                {isEditing ? (
                                    <>
                                        <button className="btn-editar" onClick={handleEditarActividad}>
                                            Guardar Cambios
                                        </button>
                                        <button
                                            className="btn-eliminar"
                                            onClick={() => {
                                                setEditandoActividad({ ...actividadSeleccionada });
                                                setIsEditing(false);
                                            }}
                                        >
                                            Cancelar
                                        </button>
                                    </>
                                ) : (
                                    <>
                                        <button className="btn-editar" onClick={() => setIsEditing(true)}>
                                            Editar Actividad
                                        </button>
                                        <button
                                            className="btn-eliminar"
                                            onClick={() => handleEliminarActividad(actividadSeleccionada.ID)}
                                        >
                                            Eliminar Actividad
                                        </button>
                                    </>
                                )}
                            </div>
                        )}
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
}

export default Activities;