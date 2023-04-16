package sort

// 归并排序
//
// 分治法+递归
func MergeSort(a []int) []int {
	n := len(a)
	if n < 2 {
		// n == 1
		return a
	}

	mid := n / 2

	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)

	res := make([]int, 0)

	leftIndex := 0
	rightIndex := 0

	for leftIndex < leftLen && rightIndex < rightLen {
		if left[leftIndex] < right[rightIndex] {
			res = append(res, left[leftIndex])
			leftIndex++
		} else {
			res = append(res, right[rightIndex])
			rightIndex++
		}
	}

	// 其中一个还没用外
	if leftIndex < leftLen {
		res = append(res, left[leftIndex:]...)
	}

	if rightIndex < rightLen {
		res = append(res, right[rightIndex:]...)
	}

	return res
}
