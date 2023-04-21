package main

import (
	"errors"
	"fmt"
	"math"
)

// 平衡二叉树节点
type AVLNode struct {
	Data   int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// 平衡二叉树节点构造函数
func NewAVLNode(data int) *AVLNode {
	return &AVLNode{
		Data:   data,
		Height: 1,
		Left:   nil,
		Right:  nil,
	}
}

// 平衡二叉树
type AVLTree struct {
	root *AVLNode
}

// 平衡二叉树构造函数
func NewAVLTree(node *AVLNode) *AVLTree {
	return &AVLTree{
		root: node,
	}
}

// 更新当前节点的高度
func (node *AVLNode) UpdataHeight() {
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	node.Height = max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 获取平衡因子
func (node *AVLNode) getBalanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	return leftHeight - rightHeight
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

// 插入
func (tree *AVLTree) Insert(data int) error {
	node, err := tree.root.Insert(data)

	tree.root = node

	return err
}

// 插入
func (node *AVLNode) Insert(data int) (newRoot *AVLNode, err error) {
	// 第一次插入
	if node == nil {
		return NewAVLNode(data), nil
	}

	// 数据已存在
	if data == node.Data {
		return node, errors.New("数据已存在")
	}

	if data < node.Data {
		// 在左边插入
		node.Left, err = node.Left.Insert(data)
	} else {
		// 在右边插入
		node.Right, err = node.Right.Insert(data)
	}

	// 更新高度
	node.UpdataHeight()

	// 旋转
	return node.adjust(), err
}

// 将node调整为平衡二叉树,并返回新的根节点
func (node *AVLNode) adjust() (newRoot *AVLNode) {
	newRoot = node
	// 获取平衡因子
	switch node.getBalanceFactor() {
	case 2:
		if node.Left.getBalanceFactor() > 0 {
			// 直接右旋
			newRoot = RightRotate(node)
		} else {
			// 先左旋,再右旋
			newRoot = LeftRightRatation(node)
		}
	case -2:
		if node.Right.getBalanceFactor() < 0 {
			// 直接左旋
			newRoot = LeftRotate(node)
		} else {
			// 先右旋,再左旋
			newRoot = RightLeftRatation(node)
		}
	}

	return newRoot
}

// 右旋
func RightRotate(node *AVLNode) *AVLNode {
	pivot := node.Left
	pivotR := pivot.Right

	// 右旋
	pivot.Right = node
	node.Left = pivotR

	// 更新高度
	node.UpdataHeight()
	pivot.UpdataHeight()

	return pivot
}

// 右旋
func LeftRotate(node *AVLNode) *AVLNode {
	pivot := node.Right
	pivotL := pivot.Left

	// 右旋
	pivot.Left = node
	node.Right = pivotL

	// 更新高度
	node.UpdataHeight()
	pivot.UpdataHeight()

	return pivot
}

// 先左旋,后右旋
func LeftRightRatation(node *AVLNode) *AVLNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
}

// 先右旋,后左旋
func RightLeftRatation(node *AVLNode) *AVLNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}

// 删除
func (tree *AVLTree) Delete(data int) error {
	node, err := tree.root.Delete(data)

	tree.root = node

	return err
}

// 删除
func (node *AVLNode) Delete(data int) (newRoot *AVLNode, err error) {
	newRoot = node

	// 直接删除
	if node == nil {
		return nil, nil
	}

	if data < node.Data {
		// 在左边
		node.Left, err = node.Left.Delete(data)

		// 删除之后更新高度
		node.UpdataHeight()
	} else if data > node.Data {
		// 在右边
		node.Right, err = node.Right.Delete(data)

		// 更新高度
		node.UpdataHeight()
	} else {
		// 找到了

		// 为叶子节点,直接删除
		if node.Left == nil && node.Right == nil {
			return nil, err
		}

		// 有两个孩子,找到右子树的最小值或左子树的最大值,在删除它
		if node.Left != nil && node.Right != nil {
			// 为了减少旋转次数,这里进行判断
			if node.getBalanceFactor() > 0 {
				// 找到左节点的最大值
				maxNode := node.Left.getMaxNode()

				node.Data = maxNode.Data

				node.Left, err = node.Left.Delete(node.Data)
			} else {
				// 找到右节点的最小值
				minNode := node.Right.getMinNode()

				node.Data = minNode.Data

				node.Right, err = node.Right.Delete(node.Data)
			}

			// 删除后更新高度
			node.UpdataHeight()

			return node, err
		}

		// 只有一个孩子,让父节点指向孩子
		if node.Left != nil {
			return node.Left, err
		}

		if node.Right != nil {
			return node.Right, err
		}
	}

	// 更新高度
	// node.UpdataHeight()

	// 调整
	return node.adjust(), err
}

// 获取以node为根节点的最小节点
func (node *AVLNode) getMinNode() *AVLNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

// 获取以node为根节点的最大节点
func (node *AVLNode) getMaxNode() *AVLNode {
	for node.Right != nil {
		node = node.Right
	}

	return node
}

// IsAVLTree 判断是不是平衡二叉树
func (tree *AVLTree) IsAVLTree() (bool, error) {
	if tree == nil || tree.root == nil {
		return true, nil
	}

	return tree.root.IsBalanced()
}

// IsBalanced 判断节点是否符合平衡二叉树的定义
func (node *AVLNode) IsBalanced() (bool, error) {
	if node == nil {
		return true, nil
	}

	// 为叶子节点
	if node.Left == nil && node.Right == nil {
		return true, nil
	}

	// 有两个节点
	if node.Left != nil && node.Right != nil {
		// 判断是否满足二叉排序树定义(node.Left.Data < node.Data && node.Data < node.Right.Data)
		if node.Left.Data > node.Data || node.Data > node.Right.Data {
			return false, fmt.Errorf("当前节点:%v, 当前节点的值:%v, 左子树数据:%v, 右子树数据:%v, 不满足二叉排序树定义", node, node.Data, node.Left.Data, node.Right.Data)
		}

		// 获取平衡因子
		bf := node.getBalanceFactor()
		if math.Abs(float64(bf)) > 1 {
			return false, fmt.Errorf("当前节点:%v, 平衡因子bf = %d", node, bf)
		}

		// 递归判断左子树和右子树
		if ok, err := node.Left.IsBalanced(); !ok {
			return ok, err
		}

		if ok, err := node.Right.IsBalanced(); !ok {
			return ok, err
		}

		return true, nil
	}

	// 此时一定只有一个节点
	if node.Left != nil {
		// 子节点必须为叶子节点
		if node.Left.Left != nil || node.Left.Right != nil {
			return false, fmt.Errorf("当前节点:%v, 非叶子节点", node)
		}

		// 子节点高度只能为1
		if node.Left.Height != 1 {
			return false, fmt.Errorf("当前节点:%v, 只有左节点, 且高度为:%d > 1 不满足", node, node.Left.Height)
		}

		// 是否满足二叉排序树
		if node.Left.Data > node.Data {
			return false, fmt.Errorf("当前节点:%v, 当前节点的值:%v, 左子树数据:%v, 不满足二叉排序树定义", node, node.Data, node.Left.Data)
		}

		return true, nil
	}

	if node.Right != nil {
		// 子节点高度只能为1
		if node.Right.Height != 1 {
			return false, fmt.Errorf("当前节点:%v, 只有左节点, 且高度为:%d > 1 不满足", node, node.Left.Height)
		}

		// 子节点必须为叶子节点
		if node.Right.Left != nil || node.Right.Right != nil {
			return false, fmt.Errorf("当前节点:%v, 非叶子节点", node)
		}

		// 是否满足二叉排序树
		if node.Right.Data < node.Data {
			return false, fmt.Errorf("当前节点:%v, 当前节点的值:%v, 左子树数据:%v, 不满足二叉排序树定义", node, node.Data, node.Left.Data)
		}

		return true, nil
	}

	return true, nil
}

func main() {
	// 平衡二叉树
	tree := NewAVLTree(nil)

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

	// 判断是否是平衡二叉树
	b, err := tree.IsAVLTree()
	fmt.Printf("是否是平衡二叉树: %v\n", b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("--------")

	err = tree.Delete(7)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder())

	// 判断是否是平衡二叉树
	b, err = tree.IsAVLTree()
	fmt.Printf("是否是平衡二叉树: %v\n", b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("--------")

	err = tree.Delete(3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder())

	// 判断是否是平衡二叉树
	b, err = tree.IsAVLTree()
	fmt.Printf("是否是平衡二叉树: %v\n", b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("--------")

	err = tree.Delete(4)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder())

	// 判断是否是平衡二叉树
	b, err = tree.IsAVLTree()
	fmt.Printf("是否是平衡二叉树: %v\n", b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Println("--------")

	err = tree.Delete(5)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	tree.PrintTree()
	fmt.Printf("Order: %v\n", tree.LevelOrder())

	// 判断是否是平衡二叉树
	b, err = tree.IsAVLTree()
	fmt.Printf("是否是平衡二叉树: %v\n", b)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
