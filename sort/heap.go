package sort

// 堆排序
func HeapSort(a []int) {
	for i := len(a); i > 1; i-- {
		// 先构建二叉堆
		buildHeap(a, i)

		// 将堆首和堆尾交换位置
		swap(&a[0], &a[i-1])
	}
}

// 构建二叉堆
func buildHeap(a []int, len int) {
	// 从第一个非叶子节点开始构建
	parent := len/2 - 1

	for parent >= 0 {
		// 调整位二叉堆
		adjustHeap(a, parent, len)

		parent--
	}
}

// 将完全二叉树调整为二叉堆
func adjustHeap(a []int, parent, len int) {
	leftIndex := 2*parent + 1
	rightIndex := 2 * (parent + 1)

	maxIndex := parent

	if rightIndex < len && a[rightIndex] > a[maxIndex] {
		maxIndex = rightIndex
	}

	if leftIndex < len && a[leftIndex] > a[maxIndex] {
		maxIndex = leftIndex
	}

	if maxIndex != parent {
		swap(&a[maxIndex], &a[parent])

		// 交换完之后可能 子节点就不是二叉堆了
		adjustHeap(a, maxIndex, len)
	}
}
