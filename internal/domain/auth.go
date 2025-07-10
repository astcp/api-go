package domain

// AuthRequest es la estructura para la solicitud de autenticación.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse es la estructura para la respuesta de autenticación.
type AuthResponse struct {
	Token string `json:"token"`
}
