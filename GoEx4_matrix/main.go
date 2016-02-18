//matrix tool is used to multiply the given matrix by its transpose matrix.
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// Exit codes.
const (
	Success = iota
	Failed
)

//issquarematrix is used to check, whether the given matrix is a square matrix or not.
func issquarematrix(size int) bool {
	root := math.Sqrt(float64(size))
	return root == math.Floor(root)
}

// transposematrix is used to form the transpose of the given matrix.
func transposematrix(matrix [10][10]int64) [10][10]int64 {
	size := int(len(matrix))
	var newmatrix [10][10]int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newmatrix[i][j] = matrix[j][i]
		}
	}
	return newmatrix
}

//multiplymatrices is used to perform the matrix multiplication.
func multiplymatrices(matrix1, matrix2 [10][10]int64, size int) [10][10]int64 {
	var solutionmatrix [10][10]int64
	var val, total int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			total = 0
			val = 0
			for k := 0; k < size; k++ {
				val = val + matrix1[i][j]*matrix2[k][j]
			}
			total = total + val
			solutionmatrix[i][j] = total
		}
	}
	return solutionmatrix
}

//printMatrix is used to print the matrix content.
func printMatrix(matrix [10][10]int64, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println()
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Input file is missing")
		os.Exit(Failed)
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(Failed)
	}
	str := string(data)
	str = strings.Trim(str, "\n")
	row := strings.Count(str, "\n") + 1
	slices := strings.Split(str, "\n")

	if !issquarematrix(row) {
		fmt.Println("not a valid matrix")
		os.Exit(Failed)
	}

	size := int(math.Sqrt(float64(row)))
	if size > 10 {
		fmt.Println("Exceeded maximum range of 10x10")
		os.Exit(Failed)
	}

	var matrix [10][10]int64

	for _, value := range slices {
		list := strings.Fields(value)
		row, _ := strconv.ParseInt(list[0], 10, 64)
		row = row - 1
		column, _ := strconv.ParseInt(list[1], 10, 64)
		column = column - 1
		matrix[row][column], _ = strconv.ParseInt(list[2], 10, 64)
	}
	fmt.Println("A = ")
	printMatrix(matrix, size)

	tmatrix := transposematrix(matrix)
	fmt.Println("A' = ")
	printMatrix(tmatrix, size)

	solution := multiplymatrices(matrix, tmatrix, size)
	fmt.Println("A * A' = ")
	printMatrix(solution, size)

}
