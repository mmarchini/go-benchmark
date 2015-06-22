package main

import "fmt"

func quickSort(list []int) []int {
	var result []int
	if len(list) == 0 {
		return result
	}
	pivot := list[0]

	var lesser []int
	var greater []int

	for i := 1; i < len(list); i++ {
		if list[i] > pivot {
			greater = append(greater, list[i])
		} else {
			lesser = append(lesser, list[i])
		}
	}

	lesser = quickSort(lesser)
	greater = quickSort(greater)

	for i := 0; i < len(lesser); i++ {
		result = append(result, lesser[i])
	}
	result = append(result, pivot)
	for i := 0; i < len(greater); i++ {
		result = append(result, greater[i])
	}

	return result
}

func main() {
	list := []int{3, 1, 2, 6, 4}
	fmt.Println("List == ", list)
	fmt.Println("List 2 == ", quickSort(list))
}
