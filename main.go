package main

import (
	"fmt"

	"github.com/maxclav/algorithms/search"
	"github.com/maxclav/algorithms/sort"
)

func main() {
	benchmark()
}

func benchmark() {
	fmt.Println("============SORT============")
	var nbSorts int = 1000
	sort.Benchmark(nbSorts)

	fmt.Println("===========SEARCH===========")
	search.Benchmark()

	fmt.Println()
}
