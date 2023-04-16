package sort

// 桶排序
func BucketSort(a []int) {
	// 桶的个数
	bucketNum := len(a)

	// 创建桶
	buckets := make([][]int, bucketNum)

	// 找到最大值
	max := getMaxInArray(a)

	// 将数据放入指定的桶
	for i := 0; i < bucketNum; i++ {
		index := (a[i] * (bucketNum - 1)) / max

		buckets[index] = append(buckets[index], a[i])
	}

	// 排序每个桶
	offset := 0
	for i := 0; i < bucketNum; i++ {
		bucketLen := len(buckets[i])

		if bucketLen > 0 {
			sortInBucket(buckets[i])

			// 将桶中的数据放入a中
			copy(a[offset:], buckets[i])

			offset += bucketLen
		}
	}
}

// 排序
func sortInBucket(a []int) {
	n := len(a)
	// 插入排序
	for i := 1; i < n; i++ {
		curValue := a[i]

		insertIndex := i

		for insertIndex > 0 {
			if curValue < a[insertIndex-1] {
				a[insertIndex] = a[insertIndex-1]
				insertIndex--
			} else {
				break
			}
		}

		a[insertIndex] = curValue
	}
}

// 获取数组中的最大值
func getMaxInArray(a []int) int {
	n := len(a)
	if n < 1 {
		return -1
	}

	max := a[0]

	for i := 1; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	return max
}
