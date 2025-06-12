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

    useEffect(() => {
        const storedLogin = localStorage.getItem('isLoggedIn') === 'true';
        const storedRole = localStorage.getItem('role');
        setIsLoggedIn(storedLogin);
        setRole(storedRole);
    }, []);

    const handleLogin = (e) => {
        e.preventDefault();

        const storedUser = localStorage.getItem('newUser');
        const storedPass = localStorage.getItem('newPass');

        if (username === 'admin' && password === 'admin') {
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('role', 'admin');
            setIsLoggedIn(true);
            setRole('admin');
            alert('Login exitoso como administrador');
            navigate('/');
        } else if (
            (username === 'user' && password === 'user') ||
            (username === storedUser && password === storedPass)
        ) {
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('role', 'user');
            setIsLoggedIn(true);
            setRole('user');
            alert('Login exitoso como usuario');
            navigate('/');
        } else {
            alert('Usuario o contraseña incorrectos');
        }
    };

    const handleRegister = (e) => {
        e.preventDefault();
        localStorage.setItem('newUser', username);
        localStorage.setItem('newPass', password);
        localStorage.setItem('newEmail', email);
        alert('Usuario registrado con éxito');
        setIsRegistering(false);
        setUsername('');
        setPassword('');
        setEmail('');
    };

    const handleLogout = () => {
        localStorage.removeItem('isLoggedIn');
        localStorage.removeItem('role');
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
                        placeholder="Usuario"
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
