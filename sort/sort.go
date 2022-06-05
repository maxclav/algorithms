package sort

import (
	"math/rand"
)

// BubbleSort sorts the slice `vals` with the
// bubble sort algorithm.
// Time: O(n^2)
// Space: O(1)
func BubbleSort(vals []int) {
	for i := 1; i < len(vals); i++ {
		for j := 1; j < len(vals)-i+1; j++ {
			if vals[j-1] > vals[j] {
				vals[j-1], vals[j] = vals[j], vals[j-1]
			}
		}
	}
}

// InsertionSort sorts the slice `vals` with the
// insertion sort algorithm.
// Time: O(n^2)
// Space: O(1)
func InsertionSort(vals []int) {
	for i := 1; i < len(vals); i++ {
		for j := i; j > 0 && vals[j-1] > vals[j]; j-- {
			vals[j-1], vals[j] = vals[j], vals[j-1]
		}
	}
}

// SelectionSort sorts the slice `vals` with the
// selection sort algorithm.
// Time: O(n^2)
// Space: O(1)
func SelectionSort(vals []int) {
	for i := 0; i < len(vals)-1; i++ {
		currentMin := i
		for j := i + 1; j < len(vals); j++ {
			if vals[j] < vals[currentMin] {
				currentMin = j
			}
		}
		if currentMin != i {
			vals[currentMin], vals[i] = vals[i], vals[currentMin]
		}
	}
}

// Quicksort sorts the slice `vals` with the
// quicksort algorithm.
// Time: O(nlogn)
// Space: O(logn)
func Quicksort(vals []int) {
	if len(vals) < 2 {
		return
	}

	// Pick a pivot (random)
	// and swap it with the value at the last index
	pivot := rand.Int() % len(vals)
	vals[pivot], vals[len(vals)-1] = vals[len(vals)-1], vals[pivot]
	pivot = len(vals) - 1

	lastSmaller := 0
	for idx, val := range vals {
		if val < vals[pivot] {
			vals[idx], vals[lastSmaller] = vals[lastSmaller], vals[idx]
			lastSmaller++
		}
	}

	vals[lastSmaller], vals[pivot] = vals[pivot], vals[lastSmaller]

	Quicksort(vals[:lastSmaller])
	Quicksort(vals[lastSmaller+1:])
}

// Mergesort sorts the slice `vals` with the
// mergesort algorithm.
// Time: O(nlogn)
// Space: O(n)
func Mergesort(vals []int) {
	res := mergesort(vals)
	copy(vals, res)
}

func mergesort(vals []int) []int {
	if len(vals) < 2 {
		return vals
	}

	middleIndex := int(len(vals) / 2)
	leftSlice, rightSlice := vals[:middleIndex], vals[middleIndex:]
	return merge(mergesort(leftSlice), mergesort(rightSlice))
}

func merge(leftSlice, rightSlice []int) []int {
	result := make([]int, 0, len(leftSlice)+len(rightSlice))

	var leftPtr, rightPtr int = 0, 0
	for leftPtr < len(leftSlice) && rightPtr < len(rightSlice) {
		leftValue, rightValue := leftSlice[leftPtr], rightSlice[rightPtr]
		if leftValue < rightValue {
			result = append(result, leftValue)
			leftPtr++
		} else {
			result = append(result, rightValue)
			rightPtr++
		}
	}

	for ; leftPtr < len(leftSlice); leftPtr++ {
		result = append(result, leftSlice[leftPtr])
	}

	for ; rightPtr < len(rightSlice); rightPtr++ {
		result = append(result, rightSlice[rightPtr])
	}

	return result
}

// Heapsort sorts the slice `vals` with the
// heapsort algorithm.
// Time: O(nlogn)
// Space: O(1)
func Heapsort(vals []int) {
	// TODO
}

// CountingSort sorts the slice `vals` with the
// counting sort algorithm.
// Time: O(n+k)
// Space: O(k)
func CountingSort(vals []int) {
	// TODO
}

// CountingSort sorts the slice `vals` with the
// counting sort algorithm.
// Time: O(n+k)
// Space: O(k)
func RadixSort(vals []int) {
	// TODO
}
