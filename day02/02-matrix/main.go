package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//1. declare matrix
	var matrix1 [3][3]int
	matrix1[0][0] = 12
	matrix1[0][1] = 15
	matrix1[0][2] = 30

	//fmt.Println(matrix1[0][1])

	//2. init array with values
	matrix2 := [5][5]int{
		{0, 1, 2, 3, 4}, // baris 0
		{1, 2, 3, 4, 5}, // baris 1
	}

	fmt.Println(matrix2[0][4])

	//displayMatrix(matrix2)

	var matrixRnadom = matrixBox(matrix2)
	displayMatrix(matrixRnadom)
	fmt.Println("")
	displayMatrix(matrixBox(matrix2))
}

func matrixBox(matrix [5][5]int) [5][5]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = rand.Intn(11)
		}
	}
	return matrix
}

func displayMatrix(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println("")
	}
}
