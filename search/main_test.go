package search_test

import (
	"fmt"
	"os"
	"search"
	"testing"
)

func TestMain(m *testing.M) {
	nums := []int{-1, 0, 3, 5, 9, 12, 13, 15, 21}

	// i := search.BinarySearch(nums, 9)
	// i := search.InterpolationSearchRecursive(nums, 9)
	i := search.InterpolationSearch(nums, 9)

	fmt.Printf("i: %v\n", i)

	os.Exit(m.Run())
}
