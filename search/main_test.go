package search_test

import (
	"fmt"
	"os"
	"search"
	"testing"
)

func TestMain(m *testing.M) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	i := search.BinarySearch(nums, 9)

	fmt.Printf("i: %v\n", i)

	os.Exit(m.Run())
}
