package http

import (
	"api-go/internal/domain"
	"api-go/internal/usecase"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler maneja las solicitudes HTTP relacionadas con la autenticación.
type AuthHandler struct {
	authUsecase usecase.AuthUsecase
	jwtProvider usecase.JWTProvider // Asumo que es el JWTProvider para generar y validar
	jwtSecret   string
}

// NewAuthHandler crea e inicializa un nuevo AuthHandler.
func NewAuthHandler(authUc usecase.AuthUsecase, jwtP usecase.JWTProvider, jwtS string) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUc,
		jwtProvider: jwtP,
		jwtSecret:   jwtS,
	}
}

// HandleLogin maneja las solicitudes de inicio de sesión.
func (h *AuthHandler) HandleLogin(c *fiber.Ctx) error {
	var req domain.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("Error al parsear el cuerpo de la solicitud de login: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.APIResponse{
			Error:   domain.ErrInvalidRequestBody.Error(),
			Details: "Por favor, proporcione un objeto JSON válido con 'username' y 'password'.",
		})
	}

	// Validar que los campos no estén vacíos ANTES de llamar al usecase
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(domain.APIResponse{
			Error:   domain.ErrInvalidRequestBody.Error(),
			Details: "Usuario y contraseña son requeridos.",
		})
	}

	token, err := h.authUsecase.Authenticate(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return c.Status(fiber.StatusUnauthorized).JSON(domain.APIResponse{
				Error:   domain.ErrInvalidCredentials.Error(),
				Details: "Credenciales inválidas",
			})
		}
		log.Printf("Error interno de autenticación para usuario '%s': %v", req.Username, err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.APIResponse{
			Error:   "Error interno del servidor",
			Details: "Fallo al autenticar: " + err.Error(),
		})
	}

	// Respuesta exitosa
	return c.Status(fiber.StatusOK).JSON(domain.APIResponse{
		Data:    domain.AuthResponse{Token: token},
		Message: "Login exitoso", // Mensaje de éxito
	})
}

// AuthMiddleware es un middleware de autenticación que valida los tokens JWT.
func (h *AuthHandler) AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.APIResponse{
			Error:   domain.ErrUnauthorized.Error(),
			Details: "El encabezado Authorization (Bearer token) está ausente.",
		})
	}

	const bearerPrefix = "Bearer "
	if !HasPrefix(authHeader, bearerPrefix) {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.APIResponse{
			Error:   domain.ErrInvalidToken.Error(),
			Details: "Se esperaba 'Bearer <token>'.",
		})
	}

	tokenString := authHeader[len(bearerPrefix):]

	err := h.jwtProvider.ValidateToken(tokenString, h.jwtSecret)
	if err != nil {
		log.Printf("Error al validar token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(domain.APIResponse{
			Error:   domain.ErrInvalidToken.Error(),
			Details: err.Error(),
		})
	}

	return c.Next()
}

// HasPrefix es una función de utilidad para verificar el prefijo (útil en gofiber)
func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}