package sort

// 冒泡排序
func BubbleSort(a []int) {
	// a := []int{9, 21, 7, 12, 33, 20, 18}
	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 将大的移动到最后
			if a[j] > a[j+1] {
				// 交换
				swap(&a[j], &a[j+1])
			}
		}
	}
}
