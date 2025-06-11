import './Login.css';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = (e) => {
        e.preventDefault();
        console.log("Intento de login:", username, password);

        // Lógica de login con roles
        if (username === 'admin' && password === 'admin') {
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('role', 'admin');
            alert('Login exitoso como administrador');
            navigate('/');
        } else if (username === 'user' && password === 'user') {
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('role', 'user');
            alert('Login exitoso como usuario');
            navigate('/');
        } else {
            alert('Usuario o contraseña incorrectos');
        }
    };

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleLogin}>
                <h1>Iniciar sesión</h1>
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
                <button type="submit">Entrar</button>
                <button
                    type="button"
                    onClick={() => navigate('/')}
                    style={{ marginTop: '10px' }}
                >
                    Volver a Home
                </button>
            </form>
        </div>
    );
};

export default Login;
