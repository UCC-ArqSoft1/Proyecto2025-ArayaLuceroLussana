import React, { useRef, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Activities from "./Activities";
import "./Home.css";

const Home = () => {
    const actividadesRef = useRef();
    const navigate = useNavigate();
    const [mostrarActividades, setMostrarActividades] = useState(false);

    const scrollToActividades = () => {
        setMostrarActividades(true);
    };

    // Cuando mostrarActividades cambia a true, hacemos scroll suavemente
    useEffect(() => {
        if (mostrarActividades && actividadesRef.current) {
            actividadesRef.current.scrollIntoView({ behavior: "smooth" });
        }
    }, [mostrarActividades]);

    return (
        <div className="body-home">
            <section className="home-hero">
                <h1 className="titulo">ALUA</h1>
                <h2 className="frase">Desafía tus límites, siente la libertad</h2>
                <div className="botones">
                    <button onClick={scrollToActividades}>Actividades</button>
                    <button onClick={() => navigate("/login")}>Iniciar sesión</button>
                </div>
            </section>

            {/* Solo renderizamos Activities si mostrarActividades es true */}
            {mostrarActividades && (
                <div ref={actividadesRef}>
                    <Activities />
                </div>
            )}
        </div>
    );
};

export default Home;
