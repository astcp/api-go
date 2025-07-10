package usecase

import (
	"fmt"

	"api-go/internal/domain"
)

// authUsecase implementa la interfaz AuthUsecase.
type authUsecase struct {
	jwtProvider JWTProvider
	jwtSecret   string
}

// NewAuthUsecase crea una nueva instancia de AuthUsecase.
func NewAuthUsecase(jwtProvider JWTProvider, jwtSecret string) AuthUsecase {
	return &authUsecase{
		jwtProvider: jwtProvider,
		jwtSecret:   jwtSecret,
	}
}

// Authenticate verifica las credenciales y genera un token JWT si son válidas.
func (uc *authUsecase) Authenticate(username, password string) (string, error) {
	if username == "admin" && password == "password" {
		token, err := uc.jwtProvider.GenerateToken(username, uc.jwtSecret)
		if err != nil {
			return "", fmt.Errorf("%w: %v", domain.ErrFailedToGenerateToken, err)
		}
		return token, nil
	}
	return "", domain.ErrInvalidCredentials
}