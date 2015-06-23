package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func matrixMultiplication(A [][]int, B [][]int) [][]int {
	M := [][]int{}
	rotB := rotateMatrix(B)

	for i := 0; i < len(A); i++ {
		row := []int{}
		for j := 0; j < len(rotB); j++ {
			row = append(row, vectorMultiplication(A[i], rotB[j]))
		}
		M = append(M, row)
	}

	return M
}

func vectorMultiplication(A []int, B []int) int {
	acc := 0
	for i := 0; i < len(A); i++ {
		acc += A[i] * B[i]
	}

	return acc
}

func matrixMultiplicationProfiler(A [][]int, B [][]int, repetitions int, filename string) {

	perf, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := perf.Close(); err != nil {
			panic(err)
		}
	}()

	for r := 0; r < repetitions; r++ {
		t := time.Now()
		result := matrixMultiplication(A, B)
		elapsed := time.Since(t).Nanoseconds()
		fmt.Println(result)

		if _, err := perf.WriteString(fmt.Sprintf("%d\n", elapsed)); err != nil {
			panic(err)
		}
	}
}

func matrixFromFile(filename string) [][]int {
	M := [][]int{}
	f, err := os.Open(filename)
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range rawCSVdata {
		row := []int{}
		for _, val := range each {
			currentValue, _ := strconv.Atoi(val)
			row = append(row, currentValue)
		}
		M = append(M, row)
	}
	return M
}

func main() {
	A := matrixFromFile(os.Args[1])
	B := matrixFromFile(os.Args[2])
	repetitions, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic(err3)
	}
	filename := os.Args[4]
	matrixMultiplicationProfiler(A, B, repetitions, filename)
}
