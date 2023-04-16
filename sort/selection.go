package sort

// 选择排序
func SelectionSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		// 找出最小的元素
		minIndex := i

		for j := i + 1; j < n; j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}

		// 移动到最前面
		if minIndex != i {
			swap(&a[minIndex], &a[i])
		}
	}
}
