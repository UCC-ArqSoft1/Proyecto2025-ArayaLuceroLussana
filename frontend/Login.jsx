import './Login.css';


const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const handleLogin = async (e) => {
        e.preventDefault();

        if (username === 'admin' && password === 'admin') {
            alert('Login successful');
            // Aquí puedes redirigir al usuario a otra página o realizar otras acciones

        } else {
            console.log('Login unsuccessful');

        }
    }

    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleLogin}>
                <h1>Enter your username</h1>
                <input
                    type="text"
                    placeholder="Usuario"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    required
                />
                <input
                    type="password"
                    placeholder="Contrasena"
                    value={password}        /*Evalua la contrasena en la base de datos donde se encuetra guardada*/
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Enter</button>
            </form>
        </div>
    );
}

createRoot(document.getElementById('root')).render(
    <React.StrictMode>
        <Login />
    </React.StrictMode>
);