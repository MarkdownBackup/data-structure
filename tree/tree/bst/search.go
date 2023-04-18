package bst

import (
	"errors"
	"fmt"
	"reflect"
)

// 二叉排序树
type BinarySearchTree struct {
	root *TreeNode
}

// 创建一个二叉排序树
func NewBinarySearchTree(root *TreeNode) *BinarySearchTree {
	return &BinarySearchTree{
		root: root,
	}
}

// 判断data的类型
func (tree *BinarySearchTree) TypeOf(data interface{}) string {
	t := reflect.TypeOf(data)
	return t.String()
}

// 插入数据
func (tree *BinarySearchTree) Insert(data interface{}) error {
	// 判断要插入的数据类型
	switch data.(type) {
	case int:
		goto int
	default:
		return fmt.Errorf("暂时不支持该类型:%v", reflect.TypeOf(data))
	}

int:
	value := data.(int)
	// 二叉排序树为空 直接插入
	node := tree.root
	if node == nil {
		tree.root = NewTreeNode(value)
		return nil
	}

	// 不为空 找到要插入的位置
	for node != nil {
		if value < node.Data.(int) {
			// 在左边

			// 判断left是否为空
			if node.Left == nil {
				// 直接插入
				node.Left = NewTreeNode(value)
				return nil
			}

			node = node.Left
		} else if value > node.Data.(int) {
			// 在右边
			if node.Right == nil {
				node.Right = NewTreeNode(value)
				return nil
			}

			node = node.Right
		} else {
			// 数据已存在 无法插入
			return fmt.Errorf("数据已存在,无法插入")
		}
	}

	return nil
}

// 中序遍历 打印整个数组
func (tree *BinarySearchTree) PrintTree() {
	res := tree.InOrder()
	fmt.Printf("order: %v\n", res)
}

// 中序遍历
func (tree *BinarySearchTree) InOrder() []interface{} {
	res := make([]interface{}, 0)

	stack := make([]*TreeNode, 0)

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
func (tree *BinarySearchTree) LevelOrder() (res [][]interface{}) {
	node := tree.root
	if node == nil {
		return
	}

	// 队列
	queue := make([]*TreeNode, 0)

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

// 查找数据为data的节点,如果不存在返回nil
func (tree *BinarySearchTree) Find(data interface{}) (*TreeNode, error) {
	// 判断数据类型
	switch data.(type) {
	case int:
		goto int
	default:
		return nil, fmt.Errorf("暂时不支持该类型:%v", reflect.TypeOf(data))
	}

int:
	value := data.(int)

	node := tree.root

	if node == nil {
		return nil, errors.New("数据不存在")
	}

	// 找到节点
	for node != nil {
		if value < node.Data.(int) {
			// 在左边
			node = node.Left
		} else if value > node.Data.(int) {
			// 在右边
			node = node.Right
		} else {
			// 找到了
			return node, nil
		}
	}

	return nil, errors.New("数据不存在")
}

func (tree *BinarySearchTree) Delete(data interface{}) error {
	// 判断数据类型
	switch data.(type) {
	case int:
		goto int
	default:
		return fmt.Errorf("暂时不支持该类型:%v", reflect.TypeOf(data))
	}

int:
	value := data.(int)

	// 要删除的节点
	node := tree.root
	// 要删除节点的父节点
	var nodeParent *TreeNode

	// 找到要删除的节点node
	for node != nil && value != node.Data.(int) {
		nodeParent = node
		if value < node.Data.(int) {
			node = node.Left
		} else if value > node.Data.(int) {
			node = node.Right
		}
	}

	if node == nil {
		return errors.New("数据不存在")
	}

	// 如果有两个节点,找到node的右子树中最小的节点代替node,并删除最小的节点
	if node.Left != nil && node.Right != nil {
		// 找到右子树的最小节点
		minNode := node.Right
		// 最小节点的父节点
		minNodeParent := node

		for minNode.Left != nil {
			minNodeParent = minNode
			minNode = minNode.Left
		}

		node.Data = minNode.Data

		// 删除最小的节点
		node = minNode
		nodeParent = minNodeParent
	}

	// 如果node只有一个子节点,只需将node的父节点指向node的子节点
	var child *TreeNode
	if node.Left != nil {
		child = node.Left
	} else if node.Right != nil {
		child = node.Right
	} else {
		// 如果node为叶子节点直接删除
		child = nil
	}

	// 删除逻辑
	if nodeParent == nil {
		// 要删除的就是根节点
		node = nil
	} else if nodeParent.Left == node {
		// 要删除的是左节点
		nodeParent.Left = child
	} else if nodeParent.Right == node {
		// 要删除的是右节点
		nodeParent.Right = child
	}

	return nil
}
