package goober

type Node[T any] struct {
	Value    []T
	Children map[rune]*Node[T]
}

func NewNode[T any]() *Node[T] {
	return &Node[T]{
		Children: make(map[rune]*Node[T], 0),
	}
}

func (n *Node[T]) getOrCreateChild(r rune) *Node[T] {
	if node, has := n.Children[r]; has {
		return node
	}

	node := NewNode[T]()
	n.Children[r] = node
	return node
}
