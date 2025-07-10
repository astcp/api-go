package domain

// QRFactorization representa el resultado de la factorización QR.
type QRFactorization struct {
	Q Matrix `json:"Q"` // Matriz ortogonal Q
	R Matrix `json:"R"` // Matriz triangular superior R
}
