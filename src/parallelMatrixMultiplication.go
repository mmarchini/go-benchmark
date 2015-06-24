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

func vectorMultiplication(A []int, B []int) int {
	acc := 0
	for i := 0; i < len(A); i++ {
		acc += A[i] * B[i]
	}

	return acc
}

func matrixMultiplication(A [][]int, B [][]int, processes int) [][]int {
	if len(A)%processes != 0 || processes > len(A) {
		panic("Erro!")
	}
	M := [][]int{}
	for i := 0; i < len(A); i++ {
		M = append(M, []int{})
	}
	rotB := rotateMatrix(B)
	c := make(chan int)
	multi := len(A) / processes

	for p := 0; p < processes; p++ {
		go func(begin int, end int) {
			for i := begin; i < end; i++ {
				for j := 0; j < len(rotB); j++ {
					M[i] = append(M[i], vectorMultiplication(A[i], rotB[j]))
				}
			}

			c <- 1
		}(p*multi, (p+1)*multi)
	}

	for p := 0; p < processes; p++ {
		<-c
	}

	return M
}

func matrixMultiplicationProfiler(A [][]int, B [][]int, processes int, repetitions int, filename string) {

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
		result := matrixMultiplication(A, B, processes)
		elapsed := time.Since(t).Nanoseconds()

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
	processes, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	repetitions, err := strconv.Atoi(os.Args[4])
	if err != nil {
		panic(err)
	}
	filename := os.Args[5]
	matrixMultiplicationProfiler(A, B, processes, repetitions, filename)
}
