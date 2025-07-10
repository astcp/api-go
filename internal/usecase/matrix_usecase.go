package usecase

import (
	"fmt"
	"math"

	"api-go/internal/domain"
	"gonum.org/v1/gonum/mat"
)

// matrixUsecase implementa la interfaz MatrixUsecase.
type matrixUsecase struct{}

// NewMatrixUsecase crea una nueva instancia de MatrixUsecase.
func NewMatrixUsecase() MatrixUsecase {
	return &matrixUsecase{}
}

// ProcessMatrix valida, rota y calcula la factorización QR de la matriz.
func (uc *matrixUsecase) ProcessMatrix(originalMatrix domain.Matrix) (rotatedMatrix domain.Matrix, qrFactorization domain.QRFactorization, err error) {
	rows := len(originalMatrix)
	if rows == 0 {
		return nil, domain.QRFactorization{}, domain.ErrMatrixEmpty
	}
	cols := len(originalMatrix[0])
	if cols == 0 {
		return nil, domain.QRFactorization{}, domain.ErrMatrixEmpty
	}
	for _, row := range originalMatrix {
		if len(row) != cols {
			return nil, domain.QRFactorization{}, domain.ErrMatrixNotRectangular
		}
	}

	rotatedMatrix = uc.rotateMatrix90Degrees(originalMatrix)

	qrFactorization, err = uc.factorizeQR(originalMatrix)
	if err != nil {
		return nil, domain.QRFactorization{}, fmt.Errorf("error al calcular la factorización QR: %w", err)
	}

	return rotatedMatrix, qrFactorization, nil
}

func (uc *matrixUsecase) rotateMatrix90Degrees(matrix domain.Matrix) domain.Matrix {
	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make(domain.Matrix, cols)
	for idx := range rotated {
		rotated[idx] = make([]float64, rows)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			rotated[c][rows-1-r] = matrix[r][c]
		}
	}
	return rotated
}

func (uc *matrixUsecase) factorizeQR(matrix domain.Matrix) (domain.QRFactorization, error) {
	rows := len(matrix)
	cols := len(matrix[0])

	flattened := make([]float64, rows*cols)
	for r := 0; r < rows; r++ {
		copy(flattened[r*cols:(r+1)*cols], matrix[r])
	}
	m := mat.NewDense(rows, cols, flattened)

	var qr mat.QR
	qr.Factorize(m)

	QMat := mat.NewDense(rows, cols, nil)
	qr.QTo(QMat)
	Q := make(domain.Matrix, rows)
	for r := 0; r < rows; r++ {
		Q[r] = make([]float64, cols)
		for c := 0; c < cols; c++ {
			Q[r][c] = QMat.At(r, c)
		}
	}

	RMat := mat.NewDense(rows, cols, nil)
	qr.RTo(RMat)
	R := make(domain.Matrix, rows)
	for r := 0; r < rows; r++ {
		R[r] = make([]float64, cols)
		for c := 0; c < cols; c++ {
			R[r][c] = RMat.At(r, c)
            if math.Abs(R[r][c]) < 1e-9 && r > c {
                R[r][c] = 0.0
            }
		}
	}

	return domain.QRFactorization{Q: Q, R: R}, nil
}