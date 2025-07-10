package usecase

import "api-go/internal/domain"

// MatrixUsecase es la interfaz para las operaciones de caso de uso de la matriz.
type MatrixUsecase interface {
	ProcessMatrix(originalMatrix domain.Matrix) (rotatedMatrix domain.Matrix, qrFactorization domain.QRFactorization, err error)
}

// AuthUsecase es la interfaz para las operaciones de caso de uso de autenticación.
type AuthUsecase interface {
	Authenticate(username, password string) (token string, err error)
}

// JWTProvider es una interfaz que la capa de Casos de Uso (o Handler) usará.
type JWTProvider interface {
	GenerateToken(username string, secret string) (string, error)
	ValidateToken(tokenString string, secret string) error
}