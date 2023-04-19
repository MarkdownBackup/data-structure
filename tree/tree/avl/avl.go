package avl

import "fmt"

// 平衡二叉树
type AVLTree struct {
	root *AVLNode // 根节点
}

// 平衡二叉树构造函数
func NewAVLTree(root *AVLNode) *AVLTree {
	return &AVLTree{
		root: root,
	}
}

// 查找
//
// 如果不存在则为nil
func (tree *AVLTree) Find(data DataType) (*AVLNode, error) {
	node := tree.root

	return node.Find(data)
}

// 插入
func (tree *AVLTree) Insert(data DataType) error {
	node, err := tree.root.Insert(data)

	tree.root = node

	return err
}

// 中序遍历 打印整个数组
func (tree *AVLTree) PrintTree() {
	res := tree.InOrder()
	fmt.Printf("order: %v\n", res)
}

// 中序遍历
func (tree *AVLTree) InOrder() []interface{} {
	res := make([]interface{}, 0)

	stack := make([]*AVLNode, 0)

	node := tree.root
	for len(stack) > 0 || node != nil {
		// 将左子树全部入栈
		for node != nil {
			stack = append(stack, node)

			node = node.Left
		}

		// 出栈
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 加入结果集
		res = append(res, node.Data)

		node = node.Right
	}

	return res
}

// 层次遍历
//
// 查看节点分布
func (tree *AVLTree) LevelOrder() (res [][]interface{}) {
	node := tree.root
	if node == nil {
		return
	}

	// 队列
	queue := make([]*AVLNode, 0)

	// 根节点入队
	queue = append(queue, node)

	for i := 0; len(queue) > 0; i++ {
		res = append(res, []interface{}{})
		n := len(queue)
		for j := 0; j < n; j++ {
			// 出队
			pop := queue[0]
			queue = queue[1:]

			res[i] = append(res[i], pop.Data)

			// 左入队
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}

			// 右入队
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
	}

	return res
}
