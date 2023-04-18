package bst

import "fmt"

// 二叉树节点
type TreeNode struct {
	Data  interface{}
	Left  *TreeNode // 左孩子
	Right *TreeNode // 右孩子
}

// 创建一个二叉树的节点
func NewTreeNode(data interface{}) *TreeNode {
	return &TreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

func (n *TreeNode) GetData() string {
	return fmt.Sprintf("%v", n.Data)
}
