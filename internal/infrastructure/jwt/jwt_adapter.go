package jwt

import (
	"errors"
	"fmt"
	"time"

	"api-go/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
)

// JWTAdapter implementa la interfaz usecase.JWTProvider.
type JWTAdapter struct{}

// NewJWTAdapter crea una nueva instancia de JWTAdapter.
func NewJWTAdapter() usecase.JWTProvider {
	return &JWTAdapter{}
}

// GenerateToken crea un nuevo token JWT.
func (a *JWTAdapter) GenerateToken(username, secret string) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user":       username,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(), // Expira en 24 horas
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("fallo al firmar el token: %w", err)
	}
	return signedToken, nil
}

// ValidateToken valida un token JWT.
func (a *JWTAdapter) ValidateToken(tokenString, secret string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return fmt.Errorf("fallo al parsear el token: %w", err)
	}

	if !token.Valid {
		return errors.New("token inválido")
	}

	return nil
}
