package domain

import "errors"

var (
	ErrMatrixEmpty         = errors.New("la matriz de entrada está vacía")
	ErrMatrixNotRectangular = errors.New("la matriz de entrada no es rectangular")
	ErrInvalidCredentials  = errors.New("credenciales inválidas")
	ErrFailedToGenerateToken = errors.New("fallo al generar el token")
	ErrUnauthorized        = errors.New("no autorizado")
	ErrInvalidToken        = errors.New("token inválido o expirado")
	ErrInvalidRequestBody  = errors.New("cuerpo de solicitud inválido")
)