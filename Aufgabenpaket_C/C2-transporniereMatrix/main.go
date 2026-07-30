package main

import "fmt"

// transposes a matrix
func transpose(matrix [5][5]int) [5][5]int {

	for i := 0; i < len(matrix)-1; i++ { // iterate over each row in the empty transposed matrix
		for j := i + 1; j < len(matrix[i]); j++ { // iterate over each row in the given matrix
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j] // add the i-te item of the row to the first row in the transposed matrix
		}
	}
	return matrix
}

// prints a matrix as a table
func printMatrix(matrix [5][5]int) {
	for _, row := range matrix {
		for _, item := range row {
			fmt.Printf("%3d", item)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	for i := 1; i < 5; i++ {
		matrix := gibMatrix(i)
		matrixT := transpose(matrix)
		fmt.Println("Die Matrix:")
		printMatrix(matrix)
		fmt.Println("wird transponiert zu:")
		printMatrix(matrixT)
	}
}
