package main

// 链表
type List[K, V interface{}] struct {
	// 头节点
	head *Node[K, V]
}

// 每一个节点
type Node[K, V interface{}] struct {
	Key   K
	Value V
	next  *Node[K, V]
}

// 创建一个链表
func New[K, V interface{}]() *List[K, V] {
	return &List[K, V]{
		head: nil,
	}
}

func NewNode[K, V interface{}](key K, value V) *Node[K, V] {
	return &Node[K, V]{
		Key:   key,
		Value: value,
		next:  nil,
	}
}

func (l *List[K, V]) init() {
	l = New[K, V]()
}

// 插入数据
func (l *List[K, V]) Insert(key K, value V) {
	// 没有初始化
	if l == nil {
		// 手动初始化
		l.init()
	}

	// 头插
	// 第一次插入
	if l.head == nil {
		l.head = NewNode(key, value)
		return
	}

	node := NewNode(key, value)
	node.next = l.head
	l.head = node
}

func (l *List[K, V]) GetAll() (k []K, v []V) {
	for i := l.head; i != nil; i = i.next {
		k = append(k, i.Key)
		v = append(v, i.Value)
	}

	return
}

func main() {
	list := New[int, int]()
	list.Insert(1, 10)
	list.Insert(2, 20)
	list.Insert(3, 30)
}
