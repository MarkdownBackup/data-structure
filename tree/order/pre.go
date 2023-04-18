package order

// 前序遍历
func PreOrderTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return
	}

	// 栈
	stack := make([]*TreeNode, 0)

	// 根节点入栈
	stack = append(stack, root)

	// 栈中还有数据
	for len(stack) > 0 {
		// 出栈
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 直接加入结果集
		res = append(res, pop.Val)

		// 右节点入栈
		if pop.Right != nil {
			stack = append(stack, pop.Right)
		}

		// 左节点入栈
		if pop.Left != nil {
			stack = append(stack, pop.Left)
		}
	}

	return res
}
