package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"api-go/internal/domain"
	
)

// Router representa el servidor HTTP y sus rutas.
type Router struct {
	app           *fiber.App
	authHandler   *AuthHandler
	matrixHandler *MatrixHandler
	port          string
}

// NewRouter crea una nueva instancia del Router.
func NewRouter(authHandler *AuthHandler, matrixHandler *MatrixHandler, port string) *Router {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router := &Router{
		app:           app,
		authHandler:   authHandler,
		matrixHandler: matrixHandler,
		port:          port,
	}

	router.setupRoutes()
	return router
}

// setupRoutes configura todas las rutas de la API.
func (r *Router) setupRoutes() {
	api := r.app.Group("/api")

	api.Post("/auth/login", r.authHandler.HandleLogin)

	// Las rutas protegidas usan el middleware de autenticación del authHandler.
	api.Post("/process-matrix", r.authHandler.AuthMiddleware, r.matrixHandler.HandleMatrixProcessing)

	r.app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(domain.APIResponse{
			Error:   "No encontrado",
			Details: "El recurso solicitado no fue encontrado. Verifique la URL.",
		})
	})
}

// Start inicia el servidor HTTP.
func (r *Router) Start() error {
	log.Printf("Go API escuchando en el puerto :%s", r.port)
	return r.app.Listen(":" + r.port)
}