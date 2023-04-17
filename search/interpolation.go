package search

// 插值查找,递归实现
func InterpolationSearchRecursive(a []int, target int) int {
	return interpolationSearchRecursive(a, target, 0, len(a)-1)
}

func interpolationSearchRecursive(a []int, target, start, end int) int {
	// 参数校验和递归终止条件
	if len(a) < 1 || start > end || start < 0 || end > len(a)-1 {
		return -1
	}

	mid := start + (end-start)*(target-a[start])/(a[end]-a[start])

	if a[mid] < target {
		return interpolationSearchRecursive(a, target, mid+1, end)
	} else if a[mid] > target {
		return interpolationSearchRecursive(a, target, start, mid-1)
	} else {
		return mid
	}
}

// 插值查找
func InterpolationSearch(a []int, target int) int {
	left, right := 0, len(a)-1

	// 参数校验
	if len(a) < 1 || target < a[left] || target > a[right] {
		return -1
	}

	for left <= right {
		mid := left + (right-left)*(target-a[left])/(a[right]-a[left])

		if target > a[mid] {
			left = mid + 1
		} else if target < a[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
