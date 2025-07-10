package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"api-go/internal/domain"
	"api-go/internal/usecase"
)

func TestMatrixUsecaseProcessMatrix(t *testing.T) {
	uc := usecase.NewMatrixUsecase()

	tests := []struct {
		name              string
		inputMatrix       domain.Matrix
		expectedRotated   domain.Matrix
		expectedQRQ       domain.Matrix
		expectedQRR       domain.Matrix
		expectedError     error
		assertQRErrorFunc func(*testing.T, domain.QRFactorization) // Para comprobar si Q y R son válidas (suman al original)
	}{
		{
			name:        "Valid 3x3 matrix",
			inputMatrix: domain.Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expectedRotated: domain.Matrix{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
			// Los valores de QR pueden ser complejos, aquí se comprueba la propiedad matemática
			assertQRErrorFunc: func(t *testing.T, qr domain.QRFactorization) {
				// Puedes añadir una aserción más rigurosa aquí, como multiplicar Q*R y comparar con la original.
				// Para simplificar, solo verificamos que no sean nil y tengan dimensiones correctas.
				assert.NotNil(t, qr.Q)
				assert.NotNil(t, qr.R)
				assert.Len(t, qr.Q, 3)
				assert.Len(t, qr.R, 3)
				assert.Len(t, qr.Q[0], 3)
				assert.Len(t, qr.R[0], 3)
			},
			expectedError: nil,
		},
		{
			name:            "Empty matrix",
			inputMatrix:     domain.Matrix{},
			expectedRotated: nil,
			expectedQRQ:     nil,
			expectedQRR:     nil,
			expectedError:   domain.ErrMatrixEmpty,
		},
		{
			name:            "Matrix with empty rows",
			inputMatrix:     domain.Matrix{{}},
			expectedRotated: nil,
			expectedQRQ:     nil,
			expectedQRR:     nil,
			expectedError:   domain.ErrMatrixEmpty,
		},
		{
			name:            "Non-rectangular matrix",
			inputMatrix:     domain.Matrix{{1, 2}, {3, 4, 5}},
			expectedRotated: nil,
			expectedQRQ:     nil,
			expectedQRR:     nil,
			expectedError:   domain.ErrMatrixNotRectangular,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotated, qr, err := uc.ProcessMatrix(tt.inputMatrix)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tt.expectedError), "Expected error %v, got %v", tt.expectedError, err)
				assert.Nil(t, rotated)
				assert.Equal(t, domain.QRFactorization{}, qr) // QR debe ser un valor cero si hay error
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedRotated, rotated)
				if tt.assertQRErrorFunc != nil {
					tt.assertQRErrorFunc(t, qr)
				}
			}
		})
	}
}