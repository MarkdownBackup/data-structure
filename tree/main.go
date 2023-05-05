package main

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

	/* // 二叉排序树
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

	fmt.Printf("tree.LevelOrder(): %v\n", tree.LevelOrder()) */

	/* // 平衡二叉树
	tree := avl.NewAVLTree(nil)

	// {3,2,1,4,5,6,7,10,9,8}
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(7)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(8)

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder())

	// a, err := tree.Find(10)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	fmt.Printf("a: %v\n", a)
	// }

	// 删除
	err := tree.Delete(3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder()) */
}
