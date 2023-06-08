package goober

import "errors"

var (
	ErrNotCompatibleKey = errors.New("the key is not compatible")
	ErrNotExists        = errors.New("the key does not exists")
)

type Storage[T any] struct {
	root  *Node[T]
	depth int
}

func NewStorage[T any](depth int) *Storage[T] {
	return &Storage[T]{
		root:  NewNode[T](),
		depth: depth,
	}
}

func (st *Storage[T]) Add(key string, value T) error {
	if len(key) != st.depth {
		return ErrNotCompatibleKey
	}

	current := st.root

	for _, r := range key {
		current = current.getOrCreateChild(r)
	}

	current.Value = append(current.Value, value)

	return nil
}

func (st *Storage[T]) Search(key string) []T {
	cursor := st.root

	for _, r := range key {
		if node, has := cursor.Children[r]; !has {
			return nil
		} else {
			cursor = node
		}

	}

	return dfs(cursor)
}

// dfs = Depth-First Search
// Depth-First search is an algorithm for traversing or searching tree or graph data structures.
func dfs[T any](node *Node[T]) []T {

	// if node.Value has values, it means that we've reached a leaf
	if len(node.Value) != 0 {
		return node.Value
	}

	var items []T

	for _, n := range node.Children {
		items = append(items, dfs(n)...)
	}

	return items
}
