package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func in(elem int, list []int) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == elem {
			return true
		}
	}
	return false
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func insertOn(solution [][]int, pos []int) [][]int {
	pos_i := pos[0]
	pos_j := pos[1]

	for s := 0; s < len(solution); s++ {
		solution_i := solution[s][0]
		solution_j := solution[s][1]
		if (abs(solution_i-pos_i) == abs(solution_j-pos_j)) || (solution_j == pos_j) || (solution_i == pos_i) {
			return [][]int{}
		}
	}

	new_solution := [][]int{}
	new_solution = append(new_solution, solution...)
	new_solution = append(new_solution, pos)

	return new_solution
}

func queensResolverAux(rows []int, columns []int, size int, current [][]int) [][][]int {
	results := [][][]int{}

	if len(columns) == size {
		ret_current := [][][]int{current}
		return ret_current
	}

	j := 0
	for in(j, columns) {
		j++
	}
	columns = append(columns, j)

	for i := 0; i < size; i++ {
		if !in(i, rows) {
			next := insertOn(current, []int{i, j})
			if len(next) > 0 {
				new_rows := append(rows, i)
				new_results := queensResolverAux(new_rows, columns, size, next)
				if len(new_results) > 0 {
					results = append(results, new_results...)
				}
			}
		}
	}

	return results
}

func queensResolverProcess(queue chan [][][]int, begin int, end int, size int) {
	results := [][][]int{}

	for i := begin; i < end; i++ {
		new_results := queensResolverAux([]int{i}, []int{0}, size, [][]int{{i, 0}})
		if len(new_results) > 0 {
			results = append(results, new_results...)
		}
	}

	queue <- results
}

func queensResolver(size int, processes int) [][][]int {
	results := [][][]int{}
	if size%processes != 0 || processes > size {
		panic("Erro!")
	}
	multi := size / processes
	queue := make(chan [][][]int)

	for p := 0; p < processes; p++ {
		go queensResolverProcess(queue, p*multi, (p+1)*multi, size)
	}

	for p := 0; p < processes; p++ {
		new_results := <-queue

		if len(new_results) > 0 {
			results = append(results, new_results...)
		}
	}

	return results
}

func queensProblemProfiler(size int, processes int, repetitions int, filename string) {

	perf, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := perf.Close(); err != nil {
			panic(err)
		}
	}()

	runtime.GOMAXPROCS(processes)

	for r := 0; r < repetitions; r++ {
		t := time.Now()
		queensResolver(size, processes)
		elapsed := time.Since(t).Nanoseconds()

		if _, err := perf.WriteString(fmt.Sprintf("%d\n", elapsed)); err != nil {
			panic(err)
		}
	}
}

func main() {
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	processes, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	repetitions, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic(err3)
	}
	filename := os.Args[4]
	queensProblemProfiler(size, processes, repetitions, filename)
}
