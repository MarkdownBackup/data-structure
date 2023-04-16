package main

import (
	"fmt"
	"my/tree/order"
)

type TreeNode = order.TreeNode

func main() {

	// 前序遍历
	// [1,null,2,3]
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	i := order.PreorderTraversal(root)
	fmt.Printf("i: %v\n", i)
}
