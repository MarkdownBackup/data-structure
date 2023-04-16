package order

// 层次遍历
func LevelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}

	queue := make([]*TreeNode, 0)

	// 根节点入队
	queue = append(queue, root)

	for i := 0; len(queue) > 0; i++ {
		// 初始化res
		res = append(res, []int{})
		n := len(queue)
		for j := 0; j < n; j++ {
			// 出队
			pop := queue[0]
			queue = queue[1:]

			res[i] = append(res[i], pop.Val)

			// 左孩子入队
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}

			// 右孩子入队
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}

		}
	}

	return res
}
