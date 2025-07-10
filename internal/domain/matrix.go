package domain

// Matrix representa una matriz rectangular.
type Matrix [][]float64

// MatrixRequest es la estructura para la entrada de la matriz en los endpoints.
type MatrixRequest struct {
	Matrix Matrix `json:"matrix"`
}