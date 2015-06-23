package main

import "fmt"

func rotateMatrix(M [][]int) [][]int {
	rotM := [][]int{}

	for i := 0; i < len(M[0]); i++ {
		row := []int{}
		for j := 0; j < len(M); j++ {
			row = append(row, M[j][i])
		}
		rotM = append(rotM, row)
	}

	return rotM
}

func vectorMultiplication(A []int, B []int) int {
	acc := 0
	for i := 0; i < len(A); i++ {
		acc += A[i] * B[i]
	}

	return acc
}

func matrixMultiplication(A [][]int, B [][]int, processes int) [][]int {
	M := [][]int{}
	rotB := rotateMatrix(B)
	c := make(chan int)
	multi := len(A) / processes

	for p := 0; p < processes; p++ {
		go func(begin int, end int) {
			for i := begin; i < end; i++ {
				row := []int{}
				for j := 0; j < len(rotB); j++ {
					row = append(row, vectorMultiplication(A[i], rotB[j]))
				}
				M = append(M, row)
			}

			c <- 1
		}(p*multi, (p+1)*multi)
	}

	for p := 0; p < processes; p++ {
		<-c
	}

	return M
}

func main() {
	A := [][]int{{0, 1, 2}, {3, 4, 5}}
	B := [][]int{{0, 1}, {2, 3}, {4, 5}}
	fmt.Println("Result:", matrixMultiplication(A, B, 2))
}
