package sort

// 计数排序
func CountingSort(a []int) {
	n := len(a)

	// 找出最大值和最小值
	max, min := getMaxAndMinInArray(a)

	// 计数数组 记录每一个元素出现的次数
	countArray := make([]int, max-min+1)

	// 初始化计数数组
	offset := -min
	for i := 0; i < n; i++ {
		countArray[a[i]+offset]++
	}

	// 遍历计数数组
	j := 0
	i := 0
	for i < n {
		if countArray[j] == 0 {
			j++
			continue
		}

		// 将数据填充到a中
		a[i] = j - offset

		countArray[j]--
		i++
	}
}

// 获取数组中最大值和最小值
func getMaxAndMinInArray(a []int) (max, min int) {
	n := len(a)
	if n < 1 {
		return -1, -1
	}

	max = a[0]
	min = a[0]

	for i := 1; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}

		if min > a[i] {
			min = a[i]
		}
	}

	return max, min
}
