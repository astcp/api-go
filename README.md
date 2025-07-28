# API Go: Procesamiento de Matrices y Autenticación JWT

## 📄 Descripción General del Proyecto

Este repositorio alberga el código fuente de una API REST desarrollada en Go, diseñada para dos propósitos principales: el procesamiento matemático de matrices y la autenticación de usuarios mediante JSON Web Tokens (JWT).

La API sigue una arquitectura de diseño limpia (inspirada en Clean Architecture), promoviendo la separación de responsabilidades y facilitando la mantenibilidad, escalabilidad y testabilidad del código.

## 🚀 Características Clave

- **Autenticación de Usuarios:** Permite a los usuarios iniciar sesión para obtener un token JWT, que es necesario para acceder a rutas protegidas.
- **Autorización Basada en JWT:** Middleware para proteger endpoints y permitir acceso solo a usuarios autenticados.
- **Procesamiento de Matrices:**
  - **Rotación:** Rota la matriz 90 grados en sentido horario.
  - **Factorización QR:** Calcula la descomposición QR (usando `gonum/matrix/mat64`).
- **Arquitectura Limpia:** Separación en capas (dominio, casos de uso, handlers, infraestructura).
- **Variables de Entorno:** Usa `.env` para gestionar configuraciones sensibles.
- **Cobertura de Pruebas Unitarias:** Pruebas para lógica de negocio y capa HTTP.

## 🛠️ Tecnologías y Dependencias

- `Go v1.22+`
- `Fiber v2`
- `joho/godotenv`
- `golang-jwt/jwt/v5`
- `stretchr/testify`
- `gonum/matrix/mat64`
- `golang.org/x/crypto/bcrypt`

## ⚙️ Configuración del Entorno de Desarrollo

### Requisitos Previos

- [Go](https://golang.org/dl/) 1.22 o superior
- [Git](https://git-scm.com/)

### Estructura del Directorio
api-go/
├── .env
├── main.go
├── go.mod
├── go.sum
├── README.md
├── internal/
│ ├── domain/
│ │ └── domain.go
│ ├── usecase/
│ │ └── usecase.go
│ ├── handler/
│ │ └── http/
│ │ ├── auth_handler.go
│ │ └── matrix_handler.go
│ └── infrastructure/
│ ├── config/
│ │ └── config.go
│ └── jwt/
│ └── jwt.go
└── tests/
├── usecase/
│ └── usecase_test.go
└── handler/
└── http/
├── auth_handler_test.go
└── matrix_handler_test.go

### ⚙️ Pasos de Configuración
1. **Crear el archivo `.env`:**

```dotenv
# .env
JWT_SECRET=tu_cadena_secreta_super_segura_para_jwt_987654321_ABCDEF
GO_API_PORT=8080
```

2. **Instalar dependencias:**

```bash
go mod tidy
```

---

### ▶️ Ejecución de la Aplicación

```bash
go run main.go
```

La API estará disponible en el puerto especificado (por defecto `:8080`).

---

### 🧪 Ejecución de Pruebas

- **Ejecutar pruebas:**

```bash
go test ./tests/usecase -v...
```
---

### 📋 Endpoints de la API

#### 1. Autenticación de Usuario

- **Endpoint:** `POST /api/auth/login`
- **Request Body:**

```json
{
  "username": "admin",
  "password": "password"
}
```

- **Respuestas:**

| Código                    | Descripción                               |
|-------------------------- |-------------------------------------------|
| 200 OK                    | Retorna JWT                               |
| 400 Bad Request           | JSON inválido o campos faltantes          |
| 401 Unauthorized          | Credenciales inválidas                    |
| 500 Internal Server Error | Error del servidor                        |

- **Ejemplo de respuesta exitosa:**

```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  },
  "message": "Login exitoso"
}
```

---

#### 2. Procesamiento de Matrices

- **Endpoint:** `POST /api/process-matrix`

- **Headers Requeridos:**
  - `Authorization: Bearer <TU_JWT_TOKEN>`
  - `Content-Type: application/json`

- **Request Body:**

```json
{
  "matrix": [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]
  ]
}
```

- **Respuestas:**

| Código | Descripción                                |
|--------|--------------------------------------------|
| 200 OK | Retorna matriz rotada y factorización QR   |
| 400 Bad Request | Matriz no válida                   |
| 401 Unauthorized | Token ausente o inválido         |
| 500 Internal Server Error | Error al procesar       |

- **Ejemplo de respuesta exitosa:**

```json
{
  "data": {
    "original_matrix": [[1, 2, 3], [4, 5, 6], [7, 8, 9]],
    "qr_factorization": {
      "Q": [[-0.13, -0.9, -0.4], [-0.5, -0.0, 0.8], [-0.8, 0.4, -0.2]],
      "R": [[-7.48, -8.76, -10.04], [0.0, -1.33, -2.66], [0.0, 0.0, 0.0]]
    },
    "rotated_matrix": [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
  },
  "message": "Matriz procesada exitosamente."
}
```

---

### 📄 Licencia

Este proyecto está bajo licencia MIT.
