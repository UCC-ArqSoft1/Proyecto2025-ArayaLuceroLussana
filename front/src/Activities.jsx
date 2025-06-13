import React, { useEffect, useState } from "react";
import "./Activities.css";

function Activities() {
    const [actividades, setActividades] = useState([]);

    useEffect(() => {
        const actividadesGuardadas = localStorage.getItem("actividadesAPI");
        if (actividadesGuardadas) {
            setActividades(JSON.parse(actividadesGuardadas));
        } else {
            fetch("http://localhost:3000/activities")
                .then((res) => res.json())
                .then((data) => setActividades(data))
                .catch((error) => {
                    console.error("Error al obtener actividades:", error);
                });
        }
    }, []);

    return (
        <div className="contenedorActividades">
            <h1>Actividades Disponibles</h1>
            <div className="listaActividades">
                {actividades.map((actividad) => (
                    <div key={actividad.id} className="actividad">
                        <h2>{actividad.nombre}</h2>
                        <p>{actividad.descripcion}</p>
                        <p><strong>Horario:</strong> {actividad.horario}</p>
                        <p><strong>Profesor:</strong> {actividad.profesor}</p>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default Activities;
