import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';  // import correcto
import './index.css';
import Login from './login.jsx';
import Home from './Home.jsx';
import Activities from './Activities.jsx';


createRoot(document.getElementById('root')).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/home" element={<Navigate to="/" replace />} />           {/* Ruta principal */}
        <Route path="/login" element={<Login />} />        {/* rutas en minúscula por convención */}
        <Route path="/activities" element={<Activities />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
