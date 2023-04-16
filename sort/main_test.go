package sort_test

import (
	"fmt"
	"my/sort"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	a := []int{9, 21, 7, 12, 33, 20, 18, 38, 28}
	// a := []int{4, 6, 2, 4, 6, 6, 7, 2}

	// sort.BubbleSort(a)
	// sort.SelectionSort(a)
	// sort.InsertionSort(a)
	// sort.QuickSort(a)
	// sort.ShellSort(a)
	// a = sort.MergeSort(a)
	// sort.HeapSort(a)
	// sort.CountingSort(a)
	// sort.BucketSort(a)
	sort.RadixSort(a)

	fmt.Printf("a: %v\n", a)

	os.Exit(m.Run())
}
