# Proyecto2025-ArayaLuceroLussana
# Enunciado Actividades Deportivas

**Arquitectura de Software I - 2025**

## Enunciado Práctico Integrado 2025

Como práctico integrador se solicita la creación de un sistema de gestión de actividades deportivas para un gimnasio, donde se destacan dos componentes:

- **Backend**, desarrollado en **Golang**, que brindará todas las interfaces necesarias para dar solución al requerimiento.
- **Frontend**, desarrollado en **React**, que implementa las vistas finales del usuario y consumirá los servicios provistos por el backend.

### Requerimientos

#### Backend

1. **Autenticación de usuarios**: 
   - Endpoints para login con manejo de sesión mediante token JWT.
   - Soporte para usuarios tipo *socio* y *administrador*.
   - Expiración del token.

2. **Funcionalidades para socios**:
   - Búsqueda de actividades disponibles.
   - Obtención por ID de una actividad.
   - Listado de actividades en las que está inscripto.
   - Inscripción a actividades (sin validación por token).

3. **Funcionalidades para administradores**:
   - Crear, editar y eliminar actividades (día, horario, cupo, categoría).
   - Validación con token.
   - Acceso a todas las funcionalidades de socios.

4. **Seguridad**:
   - Uso de token JWT firmado para escritura.
   - Contraseñas hasheadas (MD5 o SHA256).

#### Frontend

1. **Vista de inicio y búsqueda**:
   - Listado con título, horario y profesor.
   - Barra de búsqueda por palabra clave, horario o categoría.

2. **Vista de detalle**:
   - Información completa: descripción, foto, instructor, duración, etc.

3. **Inscripción**:
   - Botón para registrarse con mensaje de éxito o error.

4. **Mis actividades**:
   - Listado de actividades inscritas con acceso a detalle.

5. **Vistas de administradores**:
   - Formulario para crear actividades (solo admins).
   - Opciones para editar o eliminar desde el detalle.

---

## Consideraciones Técnicas

1. Repositorio en **GitHub**.
2. Fuente de datos: esquema **MySQL** con entidades:
   - `usuarios`, `actividades`, `inscripciones`.
3. Uso de **GORM** para el mapeo ORM.
4. Backend estructurado con patrón **MVC**.
5. Solución completa **dockerizada** (frontend, backend, DB).

---

## Condiciones de Regularidad y Examen Final

- **Regularidad**:
  - Backend: puntos 1 y 2.
  - Frontend: puntos 1, 2 y 3.
  - Técnicas: puntos 1, 2, 3 y 4.

- **Examen Final**:
  - Backend: puntos 3 y 4.
  - Frontend: puntos 4 y 5.
  - Técnicas: punto 5.

---

## Criterios de Evaluación

| Categoría | Punto | Descripción | Entrega |
|----------|-------|-------------|---------|
| Backend | Autenticación de usuarios | Login con JWT y roles | Primera Entrega |
| Backend | Seguridad (hash) | Contraseñas hasheadas | Primera Entrega |
| Backend | Funcionalidad socios | Obtener detalle por ID | Primera Entrega |
| Backend | Funcionalidad socios | Listar y buscar actividades | Primera Entrega |
| Backend | Funcionalidad socios | Listar inscripciones | Final |
| Backend | Funcionalidad socios | Inscribirse a actividad | Primera Entrega |
| Backend | Funcionalidad admins | Crear actividad (con token) | Final |
| Backend | Funcionalidad admins | Editar actividad (con token) | Final |
| Backend | Funcionalidad admins | Eliminar actividad (con token) | Final |
| Backend | Seguridad (JWT en escritura) | Operaciones protegidas | Final |
| Frontend | Vista de inicio | Muestra actividades | Primera Entrega |
| Frontend | Búsqueda | Barra funcional | Primera Entrega |
| Frontend | Detalle | Muestra información completa | Primera Entrega |
| Frontend | Inscripción | Funciona con mensaje de éxito/error | Primera Entrega |
| Frontend | Mis actividades | Listado correcto | Final |
| Frontend (Admin) | Crear actividad | Formulario funcional | Final |
| Frontend (Admin) | Editar actividad | Edición desde detalle | Final |
| Frontend (Admin) | Eliminar actividad | Eliminación desde detalle | Final |
| Base de Datos | Modelo de datos | Esquema con 3 entidades | Primera Entrega |
| Base de Datos | ORM (GORM) | Entidades mapeadas | Primera Entrega |
| Estructura Código | Patrón MVC | Organización correcta | Primera Entrega |
| DevOps | Dockerización | Todos los servicios | Final |
| DevOps | Docker Compose | Solución completa | Final |
| Control de versiones | Repositorio en GitHub | Proyecto versionado | Primera Entrega |

---

> **Nota:** Este archivo `README.md` puede incluir también instrucciones de instalación o ejecución si el repositorio las requiere.

