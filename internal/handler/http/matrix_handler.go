package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"api-go/internal/domain"
	"api-go/internal/usecase"
)

// MatrixHandler maneja las solicitudes HTTP relacionadas con el procesamiento de matrices.
type MatrixHandler struct {
	matrixUsecase usecase.MatrixUsecase
	// Ya no hay campo 'nodeClient' aquí
}

// NewMatrixHandler crea e inicializa un nuevo MatrixHandler.
func NewMatrixHandler(matrixUc usecase.MatrixUsecase) *MatrixHandler {
	return &MatrixHandler{
		matrixUsecase: matrixUc,
	}
}
// HandleMatrixProcessing maneja las solicitudes de procesamiento de matriz.
func (h *MatrixHandler) HandleMatrixProcessing(c *fiber.Ctx) error {
	var req domain.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.APIResponse{
			Error:   domain.ErrInvalidRequestBody.Error(),
			Details: "Por favor, proporcione un array de arrays de números válido en formato JSON.",
		})
	}

	rotatedMatrix, qrFactorization, err := h.matrixUsecase.ProcessMatrix(req.Matrix)
	if err != nil {
		if errors.Is(err, domain.ErrMatrixEmpty) || errors.Is(err, domain.ErrMatrixNotRectangular) {
			return c.Status(fiber.StatusBadRequest).JSON(domain.APIResponse{
				Error:   "Dimensiones de matriz inválidas",
				Details: err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(domain.APIResponse{
			Error:   "Error interno del servidor",
			Details: "Fallo al procesar la matriz: " + err.Error(),
		})
	}

	response := domain.MatrixProcessingResponse{
		OriginalMatrix:  req.Matrix,
		RotatedMatrix:   rotatedMatrix,
		QRFactorization: qrFactorization,
		// Sin campo Statistics
	}

	return c.Status(fiber.StatusOK).JSON(domain.APIResponse{
		Data:    response,
		Message: "Matriz procesada exitosamente.",
	})
}