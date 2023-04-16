package sort

import "math/rand"

// 快速排序
func QuickSort(a []int) {
	sort(a, 0, len(a)-1)
}

func sort(a []int, start, end int) {
	// 参数校验
	if len(a) < 1 || start < 0 || end > len(a)-1 || start > end {
		return
	}

	// 分区指示器
	zoneIndex := partition(a, start, end)

	// 排序分区指示器的左边
	if zoneIndex > start {
		sort(a, start, zoneIndex-1)
	}

	// 排序分区指示器的右边
	if zoneIndex < end {
		sort(a, zoneIndex+1, end)
	}
}

// 获取分区指示器
func partition(a []int, start, end int) int {
	// 参数校验
	if start == end {
		return start
	}

	// 分区指示器
	zoneIndex := start - 1

	// 随机选择一个基准数
	// [start, end+1) --> [start, end]
	pivot := rand.Intn(end-start+1) + start

	// 将基准数和最后一位交换位置
	swap(&a[pivot], &a[end])

	// 将比基准数小的值放在基准数左边 大的放在基准数右边
	for i := start; i <= end; i++ {
		// 判断当前值和[基准数]是否小于等于
		if a[i] <= a[end] {
			// 分区指示器++
			zoneIndex++

			// 分区指示器是否小于当前下标
			if zoneIndex < i {
				// 交换位置
				swap(&a[zoneIndex], &a[i])
			}
		}
	}

	return zoneIndex
}
