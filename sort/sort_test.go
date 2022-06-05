package sort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/maxclav/algorithms/utils"
	"github.com/stretchr/testify/assert"
)

// Test_BubbleSort tests BubbleSort function.
func Test_BubbleSort(t *testing.T) {
	testSortWithSizeSmall(t, BubbleSort)
}

// Test_InsertionSort test InsertionSort function.
func Test_InsertionSort(t *testing.T) {
	testSortWithSizeSmall(t, InsertionSort)
}

// Test_SelectionSort tests SelectionSort function.
func Test_SelectionSort(t *testing.T) {
	testSortWithSizeSmall(t, SelectionSort)
}

// Test_Quicksort tests Quicksort function.
func Test_Quicksort(t *testing.T) {
	testSort(t, Quicksort)
}

// Test_Mergesort tests Mergesort function.
func Test_Mergesort(t *testing.T) {
	testSort(t, Mergesort)
}

// Test_Heapsort tests Heapsort function.
func Test_Heapsort(t *testing.T) {
	// TODO
	assert.True(t, true)
}

// Test_CountingSort tests CountingSort function.
func Test_CountingSort(t *testing.T) {
	// TODO
	assert.True(t, true)
}

// Test_RadixSort tests RadixSort function.
func Test_RadixSort(t *testing.T) {
	// TODO
	assert.True(t, true)
}

func testSort(t *testing.T, sortFn func(ints []int)) {
	testSortWithSizeSmall(t, sortFn)
	testSortWithBigSmall(t, sortFn)
}

func testSortWithSizeSmall(t *testing.T, sortFn func(ints []int)) {
	for i := 0; i < 5; i++ {
		testSortWithSize(t, sortFn, i)
	}

	sliceSize := 1
	for i := 1; i < 5; i++ {
		sliceSize *= 10
		testSortWithSize(t, sortFn, sliceSize)
	}
}

func testSortWithBigSmall(t *testing.T, sortFn func(ints []int)) {
	sliceSize := 10000
	for i := 1; i < 3; i++ {
		sliceSize *= 10
		testSortWithSize(t, sortFn, sliceSize)
	}
}

func testSortWithSize(t *testing.T, sortFn func(ints []int), size int) {
	slice := utils.GenerateSlice(size)
	expected := make([]int, len(slice))
	copy(expected, slice)
	fmt.Printf("expected: %v\n", expected)
	sort.Ints(expected)
	fmt.Printf("sorted expected: %v\n", expected)
	fmt.Printf("slice: %v\n", slice)
	sortFn(slice)
	fmt.Printf("sorted slice: %v\n\n", slice)

	assert.Equal(t, expected, slice)
}
