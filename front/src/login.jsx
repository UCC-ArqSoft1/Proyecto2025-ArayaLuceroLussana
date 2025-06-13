import './Login.css';
import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [role, setRole] = useState('');
    const [isRegistering, setIsRegistering] = useState(false);

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
                const response = await fetch('http://localhost:5174/login', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        email: username,
                        password: password,
                    }),
                });

                if(!response.ok) {
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
        body: JSON.stringify({ username, password }),
            });

    if (!response.ok) {
        const error = await response.text();
        throw new Error(error || 'Error en login');
    }

    const data = await response.json();

    localStorage.setItem('isLoggedIn', 'true');
    localStorage.setItem('role', data.role || 'user'); // Ajustar según lo que devuelva tu backend

    setIsLoggedIn(true);
    setRole(data.role || 'user');
    alert('Login exitoso');
    navigate('/');
} catch (error) {
    alert(`Error al iniciar sesión: ${error.message}`);
}
    };

const handleRegister = async (e) => {
    e.preventDefault();

    try {
        const response = await fetch('http://localhost:5174/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ username, password, email }),
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error || 'Error en registro');
        }

        alert('Usuario registrado con éxito');
        setIsRegistering(false);
        setUsername('');
        setPassword('');
        setEmail('');
    } catch (error) {
        alert(`Error al registrarse: ${error.message}`);
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
            <form className="login-form" onSubmit={isRegistering ? handleRegister : handleLogin}>
                <h1>{isRegistering ? 'Crear cuenta' : 'Iniciar sesión'}</h1>

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
                {isRegistering && (
                    <input
                        type="email"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                )}
                <button type="submit">
                    {isRegistering ? 'Registrarse' : 'Entrar'}
                </button>

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
                <h2>Sesión iniciada como <strong>{role}</strong></h2>
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
