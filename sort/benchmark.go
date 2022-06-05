package sort

import (
	"fmt"
	"time"

	"github.com/maxclav/algorithms/utils"
)

func Benchmark(nbSorts int) {
	sliceSize := 1
	for i := 1; i < 5; i++ {
		sliceSize *= 10
		tests(sliceSize, nbSorts)
	}
}

func tests(sliceSize, nbSorts int) {
	fmt.Println("--------------------------")
	fmt.Printf("Size of slice: %v\n", sliceSize)

	if sliceSize <= 1000 {
		test(sliceSize, nbSorts, BubbleSort)
		test(sliceSize, nbSorts, InsertionSort)
		test(sliceSize, nbSorts, SelectionSort)
	} else {
		fmt.Printf("Size is too big for bubble, insertion and selection sort! :(\n\n")
	}
	test(sliceSize, nbSorts, Quicksort)
	test(sliceSize, nbSorts, Mergesort)
}

func test(
	sliceSize, nbSorts int,
	sortFn func(slice []int),
) {
	utils.PrintFunctionName(sortFn)
	var totalTime int64 = 0
	for i := 0; i < nbSorts; i++ {
		s := utils.GenerateSlice(sliceSize)
		timeStart := time.Now()
		sortFn(s)
		timeEnd := time.Since(timeStart).Nanoseconds()
		totalTime += timeEnd
	}
	fmt.Printf("Average: %v ns\n", int(totalTime)/nbSorts)
	fmt.Println()
}
