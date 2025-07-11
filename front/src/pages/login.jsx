import "@styles/login.css";
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import Home from './Home';

const Login = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [name, setName] = useState('');
    const [lastName, setLastName] = useState('');
    const [dni, setDni] = useState('');
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [role, setRole] = useState('');
    const [isRegistering, setIsRegistering] = useState(false);

    useEffect(() => {
        const storedLogin = localStorage.getItem('isLoggedIn') === 'true';
        const storedRole = localStorage.getItem('role');
        setIsLoggedIn(storedLogin);
        setRole(storedRole || '');
    }, []);

    const handleLogin = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, password }),
            });

            const data = await response.json();

            if (!response.ok) {
                console.error("Error en login:", data.message);
                alert(data.message || "Error al iniciar sesión");
                return;
            }

            const token = data.token;

            localStorage.setItem("token", token);
            localStorage.setItem("isLoggedIn", "true");

            // Decodificar el payload del JWT para obtener rol y userId
            const payloadBase64 = token.split('.')[1];
            const decodedPayload = JSON.parse(atob(payloadBase64));
            const userRole = decodedPayload.rol || decodedPayload.role || "";
            const userId = decodedPayload.id || decodedPayload.userId || decodedPayload.ID || null;

            if (!userId) {
                alert("Error: no se pudo obtener el userId del token.");
                return;
            }

            localStorage.setItem("role", userRole);
            localStorage.setItem("userId", userId.toString());

            setIsLoggedIn(true);
            setRole(userRole);
            navigate("/home");
        } catch (error) {
            console.error("Error en login:", error);
        }
    };

    const handleRegister = async (e) => {
        e.preventDefault();

        if (!name || !lastName || !dni || !email || !password) {
            alert('Por favor completa todos los campos');
            return;
        }

        try {
            const response = await fetch('http://localhost:8080/register', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name,
                    lastName,
                    DNI: Number(dni), // El backend espera un número
                    email,
                    password,
                    rol: 'Socio', // rol por defecto
                }),
            });

            const data = await response.json();

            if (!response.ok) {
                alert(`Error en registro: ${data.message || JSON.stringify(data)}`);
                return;
            }

            alert('Usuario registrado con éxito, ya podés iniciar sesión.');
            setIsRegistering(false);
            setName('');
            setLastName('');
            setDni('');
            setEmail('');
            setPassword('');
        } catch (error) {
            console.error('Error al registrarse:', error);
            alert('Error de conexión con el servidor');
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('isLoggedIn');
        localStorage.removeItem('role');
        localStorage.removeItem('token');
        localStorage.removeItem('userId');

        setIsLoggedIn(false);
        setRole('');
        alert('Sesión cerrada correctamente');
        navigate('/');
    };

    return (
        <div className="login-container">
            {!isLoggedIn ? (
                <form
                    className="login-form"
                    onSubmit={isRegistering ? handleRegister : handleLogin}
                >
                    <h1>{isRegistering ? 'Crear cuenta' : 'Iniciar sesión'}</h1>

                    {isRegistering && (
                        <>
                            <input
                                type="text"
                                placeholder="Nombre"
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                required
                            />
                            <input
                                type="text"
                                placeholder="Apellido"
                                value={lastName}
                                onChange={(e) => setLastName(e.target.value)}
                                required
                            />
                            <input
                                type="number"
                                placeholder="DNI"
                                value={dni}
                                onChange={(e) => setDni(e.target.value)}
                                required
                            />
                        </>
                    )}

                    <input
                        type="email"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />

                    <button type="submit">{isRegistering ? 'Registrarse' : 'Entrar'}</button>

                    <button
                        type="button"
                        className="btn-home"
                        onClick={() => navigate('/')}
                    >
                        Volver a Home
                    </button>

                    <button
                        type="button"
                        className="btn-home"
                        onClick={() => setIsRegistering(!isRegistering)}
                    >
                        {isRegistering
                            ? '¿Ya tenés cuenta? Iniciar sesión'
                            : '¿No tenés cuenta? Registrate'}
                    </button>
                </form>
            ) : (
                <div className="logout-section">
                    <h2>
                        Sesión iniciada como <strong>{role}</strong>
                    </h2>
                    <button className="btn-logout" onClick={handleLogout}>
                        Cerrar sesión
                    </button>
                    <button
                        className="btn-home"
                        onClick={() => navigate('/')}
                        style={{ marginTop: '1rem' }}
                    >
                        Ir a Home
                    </button>
                </div>
            )}
        </div>
    );
};

export default Login;
