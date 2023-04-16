package sort

// 基数排序
func RadixSort(a []int) {
	// a := []int{9, 21, 7, 12, 33, 20, 18, 38, 28}
	n := len(a)

	// 基数，十进制10、二进制2、byte型字符串256等
	radix := 10

	// 获取最大值，意味着循环几轮
	max := getMaxInArray(a)

	round := 1

	for max > 0 {
		// 创建桶
		buckets := make([][]int, radix)

		// 初始化桶
		for i := 0; i < n; i++ {
			index := (a[i] % (radix * round)) / round

			buckets[index] = append(buckets[index], a[i])
		}

		// 将桶中的数据copy到数组
		offset := 0
		for i := 0; i < radix; i++ {
			bucketlen := len(buckets[i])

			if bucketlen > 0 {
				copy(a[offset:], buckets[i])

				offset += bucketlen
			}
		}

		max /= 10
		round *= 10
	}
}
