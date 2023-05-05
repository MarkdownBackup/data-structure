package redblacktree

// 当前迭代器所在位置
type position byte

const (
	begin   position = 0
	between position = 1
	end     position = 2
)

// 迭代器
type Iterator struct {
	tree     *Tree
	node     *Node
	position position
}

func (iterator *Iterator) Next() bool {

	return false
}
