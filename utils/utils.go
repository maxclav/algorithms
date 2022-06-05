package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

// GenerateSlice generates a slice.
func GenerateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// PrintFunctionName prints the function name.
func PrintFunctionName(i interface{}) {
	fmt.Println(GetFunctionName(i))
}

// GetFunctionName returns the function name
// in string format.
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
