package redblacktree

// 颜色
type color bool

const (
	red   color = false
	black color = true
)

// 红黑树节点
type Node struct {
	Key    interface{}
	Value  interface{}
	Left   *Node // 左节点
	Right  *Node // 右节点
	Parent *Node // 父节点
	color  color // 当前节点颜色
}

// 返回元素的总个数。
// 在每次调用时动态计算，即遍历子树以计算节点数。
func (node *Node) Size() int {
	if node == nil {
		return 0
	}

	size := 1

	// 左节点还有数据
	if node.Left != nil {
		size += node.Left.Size()
	}

	// 右节点还有数据
	if node.Right != nil {
		size += node.Right.Size()
	}

	return size
}

// 祖父节点
func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil && node.Parent.Parent != nil {
		return node.Parent.Parent
	}

	return nil
}

// 兄弟节点
func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}

	// 此时 node != nil && node.Parent != nil
	if node == node.Parent.Left {
		return node.Parent.Right
	}

	return node.Parent.Left
}

// 叔伯节点(祖父节点的孩子)
func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}

	// 父节点的兄弟就是叔伯
	return node.Parent.sibling()
}

// 获取以当前节点为根节点的最大节点
func (node *Node) maxNode() *Node {
	if node == nil {
		return nil
	}

	for node.Right != nil {
		node = node.Right
	}

	return node
}

// 获取当前节点的颜色
func (node *Node) getColor() color {
	// 根据红黑树定义,叶子节点全都为黑色
	if node == nil {
		return black
	}

	return node.color
}
