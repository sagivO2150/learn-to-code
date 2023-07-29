package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{3, 1, 4, 2}
	fmt.Println("before sorting -", a)
	fmt.Println("using the sort function", sortThis(a))
}

func sortThis(x []float64) []float64 {
	y := make([]float64, len(x))
	sort.Float64s(x)
	copy(y, x)

	return y
}
