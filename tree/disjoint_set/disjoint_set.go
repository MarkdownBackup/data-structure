package disjointset

type ElementType = byte

// 并查集
type UnionSet struct {
	Data   []ElementType
	Father []int
	len    int
}

func NewUnionSet(n int) UnionSet {
	father := make([]int, n)
	data := make([]ElementType, n)

	// 初始化
	for i := 0; i < n; i++ {
		father[i] = -1
	}

	us := UnionSet{
		Father: father,
		Data:   data,
		len:    n,
	}

	return us
}

// 根据当前数据查找根节点
func (us *UnionSet) Find(x ElementType) int {
	return -1
}

func main() {
	// us := NewUnionSet(10)
}
