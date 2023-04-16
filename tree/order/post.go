package order

// 后续遍历
func postorderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)

	// 根节点入栈
	stack = append(stack, root)

	for len(stack) > 0 {
		// 出栈
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 头插法插入
		res = append([]int{pop.Val}, res...)

		// 左孩子入栈
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}

		// 右孩子入栈
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}
	}

	return res
}
