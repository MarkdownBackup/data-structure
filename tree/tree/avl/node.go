package avl

import "errors"

// 节点存储的数据类型
type DataType = int

// 二叉树节点
type AVLNode struct {
	Data   DataType
	Left   *AVLNode // 左孩子
	Right  *AVLNode // 右孩子
	Height int      // 当前节点的高度,用于计算平衡因子
}

// 创建一个二叉树的节点
func NewAVLNode(data DataType) *AVLNode {
	return &AVLNode{
		Data:   data,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}

// 查找
//
// 如果不存在则为nil
func (node *AVLNode) Find(data DataType) (*AVLNode, error) {
	for node != nil {
		if data > node.Data {
			// 在右边
			node = node.Right
		} else if data < node.Data {
			// 在左边
			node = node.Left
		} else {
			// 找到了
			return node, nil
		}
	}

	return nil, errors.New("节点不存在")
}

// 插入
//
// 返回插入后的新根节点(如果根节点变化的话)
func (node *AVLNode) Insert(data DataType) (*AVLNode, error) {
	// 节点不存在
	if node == nil {
		// 直接插入
		return NewAVLNode(data), nil
	}

	// 如果值重复,插入失败
	if node.Data == data {
		return node, errors.New("值已存在,插入失败")
	}

	// 返回的错误信息，初始化为nil
	var err error

	// 新的根节点
	var newNodeTree *AVLNode

	// 判断要插入的位置
	if data < node.Data {
		// 要插入左边
		node.Left, err = node.Left.Insert(data)
		// 计算平衡因子
		// 平衡二叉树特性:平衡因子不能大于1
		bf := node.getBalanceFactor()
		if bf == 2 {
			// 插入到了node左子树的左孩子
			if data < node.Left.Data {
				// 直接右旋
				newNodeTree = RightRotate(node)
			} else {
				// 插入到了node左子树的右孩子
				// 先左旋，后右旋
				newNodeTree = LeftRightRotation(node)
			}
		}
	} else {
		// 插入右边
		node.Right, err = node.Right.Insert(data)
		// 计算平衡因子
		// 平衡二叉树特性:平衡因子不能大于1
		bf := node.getBalanceFactor()
		if bf == -2 {
			// 插入到了node右子树的左孩子
			if data > node.Right.Data {
				// 直接左旋
				newNodeTree = LeftRotate(node)
			} else {
				// 插入到了node右子树的左孩子
				// 先右旋，后左旋
				newNodeTree = RightLeftRotation(node)
			}
		}
	}

	// 说明需要调整根节点
	if newNodeTree != nil {
		// 更新节点的高度
		newNodeTree.UpdataHeight()
		return newNodeTree, err
	}

	// 此时 newNodeTree == nil
	// 说明插入节点后,整棵树还是平衡的,不需要修改根节点
	node.UpdataHeight()
	return node, err
}

// 计算当前节点的平衡因子
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

// 更新高度
func (node *AVLNode) UpdataHeight() {
	if node == nil {
		return
	}

	// 分别计算左高度和右高度
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}

	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 更新为大的
	maxHeight := leftHeight
	if maxHeight < rightHeight {
		maxHeight = rightHeight
	}

	node.Height = maxHeight + 1
}

// 右旋
func RightRotate(node *AVLNode) *AVLNode {
	/* 树结构示意图：
	       y                x
	      / \      右旋    / \
	     x   O    ----->  O   y
	    / \ 	               / \
	   O   r                r   O
	*/

	// 调整
	pivot := node.Left
	pivotR := pivot.Right

	pivot.Right = node
	node.Left = pivotR

	// 更新高度
	node.UpdataHeight()
	pivot.UpdataHeight()

	return pivot
}

// 左旋
func LeftRotate(node *AVLNode) *AVLNode {
	/* 树结构示意图：
	       y                x
	      / \      左旋    / \
	     x   O    <-----  O   y
	    / \ 	               / \
	   O   l                l   O
	*/

	// 调整
	pivot := node.Right
	pivotL := pivot.Left

	pivot.Left = node
	node.Right = pivotL

	// 更新高度
	node.UpdataHeight()
	pivot.UpdataHeight()

	return pivot
}

// 双旋操作（先左后右）
func LeftRightRotation(node *AVLNode) *AVLNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
}

// 先右旋后左旋
func RightLeftRotation(node *AVLNode) *AVLNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}
