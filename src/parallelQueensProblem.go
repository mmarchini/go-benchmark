package main

import (
	"fmt"
	"os"
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

func testSolEq(sol1 [][]int, sol2 [][]int) bool {
	for i := 0; i < len(sol1); i++ {
		if sol1[i][0] != sol2[i][0] || sol1[i][1] != sol2[i][1] {
			return false
		}
	}

	return true
}

func testSolIn(sol [][]int, solutions [][][]int) bool {
	for s := 0; s < len(solutions); s++ {
		if testSolEq(sol, solutions[s]) {
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

	insert_pos := 0
	for s := 0; s < len(solution); s++ {
		solution_i := solution[s][0]
		solution_j := solution[s][1]
		if (abs(solution_i-pos_i) == abs(solution_j-pos_j)) || (solution_j == pos_j) || (solution_i == pos_i) {
			return [][]int{}
		}
		if solution_i < pos_i {
			insert_pos = s + 1
		}
	}

	new_solution := [][]int{}
	if insert_pos == 0 {
		new_solution = append(new_solution, pos)
		new_solution = append(new_solution, solution...)
	} else if insert_pos == len(solution) {
		new_solution = append(new_solution, solution...)
		new_solution = append(new_solution, pos)
	} else {
		new_solution = append(new_solution, solution[:insert_pos]...)
		new_solution = append(new_solution, pos)
		new_solution = append(new_solution, solution[insert_pos:]...)
	}

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
					for r := 0; r < len(new_results); r++ {
						if !testSolIn(new_results[r], results) {
							results = append(results, new_results[r])
						}
					}
				}
			}
		}
	}

	return results
}

func queensResolverProcess(queue chan [][][]int, begin int, end int, size int) {
	results := [][][]int{}

	for j := begin; j < end; j++ { //
		for i := 0; i < size; i++ {
			new_results := queensResolverAux([]int{i}, []int{j}, size, [][]int{{i, j}})
			if len(new_results) > 0 {
				for r := 0; r < len(new_results); r++ {
					if !testSolIn(new_results[r], results) {
						results = append(results, new_results[r])
					}
				}
			}
		}
	}

	queue <- results
}

func queensResolver(size int, processes int) [][][]int {
	results := [][][]int{}
	if size&processes != 0 {
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
			for r := 0; r < len(new_results); r++ {
				if !testSolIn(new_results[r], results) {
					results = append(results, new_results[r])
				}
			}
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
