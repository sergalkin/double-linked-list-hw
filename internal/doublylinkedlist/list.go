package doublylinkedlist

type List struct {
	root Node
	len  int
}

func New() *List {
	return new(List).Init()
}

/**
 Создаем sentinel node, которая будет являться указателем на первую и последнюю node списка
 */
func (l *List) Init() *List {
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

func (l *List) Len() int {
	return l.len
}

func (l *List) First() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Last() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) Push(value interface{}) {
	node := Node{
		Value: value,
	}

	if l.len == 0 {
		l.root.next = &node
	} else {
		prevTail := l.root.prev
		node.prev = prevTail
		prevTail.next = &node
	}
	l.len++
	l.root.prev = &node
}
