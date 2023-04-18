package order

func InOrderTraversal(root *TreeNode) (res []int) {
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		// 将左节点全部入栈
		for root != nil {
			stack = append(stack, root)

			root = root.Left
		}

		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 加入结果集
		res = append(res, root.Val)

		root = root.Right
	}

	return res
}
