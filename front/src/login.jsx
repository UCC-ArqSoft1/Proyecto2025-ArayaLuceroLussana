import './Login.css';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [role, setRole] = useState('');

    // Al cargar el componente, recuperamos el estado de sesión
    useEffect(() => {
        const storedLogin = localStorage.getItem('isLoggedIn') === 'true';
        const storedRole = localStorage.getItem('role');
        setIsLoggedIn(storedLogin);
        setRole(storedRole || '');
    }, []);

    const handleLogin = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email: username,
                    password: password,
                }),
            });

            if (!response.ok) {
                const err = await response.json();
                alert(err.message || 'Error al iniciar sesión');
                return;
            }

            const data = await response.json();
            const token = data.Token;

            // Decodificamos el JWT para obtener el rol del usuario
            const base64Url = token.split('.')[1];
            const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
            const payload = JSON.parse(atob(base64));
            const userRole = payload.Rol;

            // Guardamos los datos en localStorage
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('role', userRole);
            localStorage.setItem('token', token);

            setIsLoggedIn(true);
            setRole(userRole);
            alert(`Login exitoso como ${userRole}`);
            navigate('/');
        } catch (error) {
            console.error('Error en login:', error);
            alert('Error de conexión con el servidor');
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('isLoggedIn');
        localStorage.removeItem('role');
        localStorage.removeItem('token');
        setIsLoggedIn(false);
        setRole('');
        alert('Sesión cerrada correctamente');
        navigate('/');
    };

    return (
        <div className="login-container">
            {!isLoggedIn ? (
                <form className="login-form" onSubmit={handleLogin}>
                    <h1>Iniciar sesión</h1>
                    <input
                        type="text"
                        placeholder="Usuario (email)"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="Contraseña"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <button type="submit">Entrar</button>
                    <button
                        type="button"
                        className="btn-home"
                        onClick={() => navigate('/')}
                    >
                        Volver a Home
                    </button>
                </form>
            ) : (
                <div className="logout-section">
                    <h2>Ya has iniciado sesión como <strong>{role}</strong></h2>
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
