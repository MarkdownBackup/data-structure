package main

import (
	"fmt"
	"my/tree/order"
	"my/tree/tree/bst"
)

type TreeNode = order.TreeNode

type test struct {
	data int
}

func main() {

	// 前序遍历
	// [1,null,2,3]
	// root := &TreeNode{
	// 	Val: 1,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val: 3,
	// 		},
	// 	},
	// }
	//
	// i := order.LevelOrder(root)
	// fmt.Printf("i: %v\n", i)

	// 二叉排序树
	tree := bst.NewBinarySearchTree(nil)

	// 插入
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)

	tree.PrintTree()

	fmt.Printf("tree.LevelOrder(): %v\n", tree.LevelOrder())

	// 查找
	// node, err := tree.Find(3)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Printf("node: %v\n", node)
	// }

	// 删除
	fmt.Println("删除值为 3 的节点：")
	err := tree.Delete(3)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("删除成功")
	}
	tree.PrintTree()

	fmt.Printf("tree.LevelOrder(): %v\n", tree.LevelOrder())
}
