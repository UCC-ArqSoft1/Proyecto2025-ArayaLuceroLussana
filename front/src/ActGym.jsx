import { useEffect, useState } from 'react';

const Act = () => {
    const [act, setAct] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/actgym') //Es como hacer un GET
            .then((res) => res.json()) //Guardamos respuesta en forma de json
            .then((data) => setAct(data)) //Guardamos en variable vinilos
            .catch((err) => console.error('Error fetching activities:', err)) //En caso de error, larga el mensaje
    }, []);



    return (
        <div className="act-container">
            {act.map((activity, index) => (
                <div className="car-act" key={index}>
                    <h2>{act.id} </h2>
                    <h2>{act.nombre} </h2>
                    <h2>{act.descripcion} </h2>
                    <h2>{act.profesor} </h2>
                    <h2>{act.horarios} </h2>
                    <h2>{act.cupos} </h2>
                </div>
            ))}
        </div>
    )
}
export default Act;