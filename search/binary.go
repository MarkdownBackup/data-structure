package search

// 二分查找(折半查找)
//
// 没找到就返回-1
func BinarySearch(a []int, target int) int {
	n := len(a)
	start, end := 0, n-1

	for start <= end {
		mid := (end + start) / 2
		if a[mid] < target {
			start = mid + 1
		} else if a[mid] > target {
			end = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
