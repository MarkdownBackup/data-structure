package search_test

import (
	"fmt"
	"os"
	"search"
	"testing"
)

func TestMain(m *testing.M) {
	a := []int{-1, 0, 3, 5, 9, 12, 13, 15, 21}
	// a := []int{2}

	// i := search.BinarySearch(a, 9)
	// i := search.InterpolationSearchRecursive(a, 9)
	// i := search.InterpolationSearch(a, 9)
	i := search.FibonacciSearch(a, 12)

	fmt.Printf("i: %v\n", i)

	os.Exit(m.Run())
}
