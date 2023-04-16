package sort

// 希尔排序
//
// 是对插入排序的一种优化
func ShellSort(a []int) {
	n := len(a)

	// 增量
	gap := n / 2

	for gap > 0 {

		// 默认gap前有序
		for i := gap; i < n; i++ {
			// 当前元素的值
			curValue := a[i]

			// 待插入位置的下标
			insertIndex := i

			// 遍历已排好序的数组
			for insertIndex >= gap {
				if curValue < a[insertIndex-gap] {
					a[insertIndex] = a[insertIndex-gap]
				} else {
					// HACK:这里总是忘记!!!!
					break
				}

				insertIndex -= gap
			}

			// 插入
			a[insertIndex] = curValue
		}

		gap /= 2
	}
}
