package search

// 斐波那契查找(黄金分割查找)
func FibonacciSearch(a []int, target int) int {
	n := len(a)

	// 构建斐波那契数列
	fib := buildFib(n)

	k := 0
	// 找到斐波那契数列中与n最近的数k(比n第一个大的数)
	for fib[k] < n {
		k++
	}

	// 临时数组
	temp := make([]int, fib[k]-n)
	a = append(a, temp...)

	// 将a扩大到k那么大
	for i := n; i < fib[k]; i++ {
		a[i] = a[n-1]
	}

	left := 0
	right := n - 1
	mid := 0
	for left <= right {
		// 找到mid
		if k == 0 {
			// 说明只有一个元素
			mid = left
		} else {
			mid = left + fib[k-1] - 1
		}

		if a[mid] > target {
			// 目标值在左边
			right = mid - 1
			k--
		} else if a[mid] < target {
			// 在右边
			left = mid + 1
			k -= 2
		} else {
			// 判断在没在范围内
			if mid <= right {
				// 找到了 就是mid
				return mid
			} else {
				// 越界了 返回最后一个元素
				return right
			}
		}
	}

	return -1
}

// 构建长度位n的斐波那契数列
func buildFib(n int) []int {
	if n < 10 {
		n = 10
	}

	fib := make([]int, n)
	fib[0] = 1
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
