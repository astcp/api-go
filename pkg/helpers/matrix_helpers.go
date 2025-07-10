package helpers

import (
	"fmt"
	"strings"

	"api-go/internal/domain"
)

// PrintMatrix imprime una matriz en formato legible.
func PrintMatrix(m domain.Matrix, name string) {
	fmt.Printf("--- %s ---\n", name)
	for _, row := range m {
		var rowStrings []string
		for _, val := range row {
			rowStrings = append(rowStrings, fmt.Sprintf("%.2f", val))
		}
		fmt.Printf("[%s]\n", strings.Join(rowStrings, ", "))
	}
	fmt.Println("-------------")
}
