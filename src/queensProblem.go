package main

import "fmt"

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

func remove(list []int, elem int) []int {
	for i := 0; i < len(list); i++ {
		if list[i] == elem {
			list = append(list[:i], list[i+1])
		}
	}

	return list
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

func queensResolver(size int) [][][]int {
	results := [][][]int{}

	for j := 0; j < size; j++ { //
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

	return results
}

func main() {
	fmt.Println("Resolution count (1):", len(queensResolver(1)))
	fmt.Println("Resolution count (2):", len(queensResolver(2)))
	fmt.Println("Resolution count (3):", len(queensResolver(3)))
	fmt.Println("Resolution count (4):", len(queensResolver(4)))
	fmt.Println("Resolution count (5):", len(queensResolver(5)))
	fmt.Println("Resolution count (6):", len(queensResolver(6)))
	fmt.Println("Resolution count (7):", len(queensResolver(7)))
	fmt.Println("Resolution count (8):", len(queensResolver(8)))
	fmt.Println("Resolution count (9):", len(queensResolver(9)))
	fmt.Println("Resolution count (10):", len(queensResolver(10)))
	fmt.Println("Resolution count (11):", len(queensResolver(11)))
	fmt.Println("Resolution count (12):", len(queensResolver(12)))
	fmt.Println("Resolution count (13):", len(queensResolver(13)))
}
