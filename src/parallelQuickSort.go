package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

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

	lesser_channel := make(chan int)
	go func() {
		lesser = quickSort(lesser)
		lesser_channel <- 1
	}()
	greater_channel := make(chan int)
	go func() {
		greater = quickSort(greater)
		greater_channel <- 1
	}()

	<-lesser_channel
	for i := 0; i < len(lesser); i++ {
		result = append(result, lesser[i])
	}
	result = append(result, pivot)
	<-greater_channel
	for i := 0; i < len(greater); i++ {
		result = append(result, greater[i])
	}

	return result
}

func quickSortProfiler(list []int, processes int, repetitions int, filename string) {

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
		quickSort(list)
		elapsed := time.Since(t).Nanoseconds()

		if _, err := perf.WriteString(fmt.Sprintf("%d\n", elapsed)); err != nil {
			panic(err)
		}
	}
}

func listFromFile(filename string) []int {
	list := []int{}
	f, err := os.Open(filename)

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, each := range rawCSVdata {
		currentValue, _ := strconv.Atoi(each[0])
		list = append(list, currentValue)
	}
	return list
}

func main() {
	problem_file := os.Args[1]
	processes, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		panic(err2)
	}
	repetitions, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic(err3)
	}
	filename := os.Args[3]
	list := listFromFile(problem_file)
	quickSortProfiler(list, processes, repetitions, filename)
}
