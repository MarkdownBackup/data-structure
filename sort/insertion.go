package sort

// 插入排序
func InsertionSort(a []int) {
	n := len(a)
	// 默认第一个有序
	for i := 1; i < n; i++ {
		// 当前元素的值
		curValue := a[i]

		// 要插入的下标位置
		insertIndex := i

		// 逆序遍历已排好序的数组下标
		for insertIndex > 0 {
			if curValue < a[insertIndex-1] {
				a[insertIndex] = a[insertIndex-1]
				insertIndex--
			} else {
				break
			}
		}

		// 差入
		a[insertIndex] = curValue
	}
}
