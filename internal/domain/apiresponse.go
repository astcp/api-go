package domain

// APIResponse es la estructura de respuesta genérica para los endpoints.
type APIResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Details string      `json:"details,omitempty"`
}

// MatrixProcessingResponse es la estructura de respuesta específica para el procesamiento de matrices.
type MatrixProcessingResponse struct {
	OriginalMatrix  Matrix          `json:"original_matrix"`
	RotatedMatrix   Matrix          `json:"rotated_matrix,omitempty"`
	QRFactorization QRFactorization `json:"qr_factorization"`
}
