package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig contiene la configuración de la aplicación.
type AppConfig struct {
	JWTSecret string
	APIPort    string
}

// LoadConfig carga las variables de entorno en la estructura AppConfig.
func LoadConfig() (*AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Advertencia: No se pudo cargar el archivo .env, usando variables de entorno del sistema: %v\n", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET no configurado en .env o variables de entorno")
	}
	
	apiPort := os.Getenv("GO_API_PORT")
	if apiPort == "" {
		apiPort = "8080" // Valor por defecto
	}

	return &AppConfig{
		APIPort:    apiPort,
	}, nil
}