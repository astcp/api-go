package main

import (
	"log"

	"api-go/internal/handler/http"
	"api-go/internal/infrastructure/config"
	"api-go/internal/infrastructure/jwt"
	"api-go/internal/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	// Inicializar adaptadores de infraestructura (implementaciones de puertos).
	jwtAdapter := jwt.NewJWTAdapter() // Implementación de usecase.JWTProvider

	// Inicializar casos de uso (core de la lógica de negocio).
	matrixUsecase := usecase.NewMatrixUsecase()
	authUsecase := usecase.NewAuthUsecase(jwtAdapter, cfg.JWTSecret)

	// Inicializar handlers HTTP (capa de entrada/presentación).
	authHandler := http.NewAuthHandler(authUsecase, jwtAdapter, cfg.JWTSecret)
	// matrixHandler ya no necesita el cliente de Node.js
	matrixHandler := http.NewMatrixHandler(matrixUsecase)

	// Inicializar Router y Servidor HTTP.
	router := http.NewRouter(authHandler, matrixHandler, cfg.APIPort)

	if err := router.Start(); err != nil {
		log.Fatalf("Fallo al iniciar el servidor Go API: %v", err)
	}
}