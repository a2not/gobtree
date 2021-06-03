package gobtree

type Item interface {
	Less(other Item) bool
}

type node struct {
	items    []Item
	children []*node
}

type BTree struct {
	degree int
	size   int // number of elements in a whole tree
	root   *node
}

func New(degree int) *BTree {
	return &BTree{
		degree: degree,
	}
}

// maximum number of items in a node
func (t *BTree) maxItems() int {
	return t.degree*2 - 1
}

// minimum number of items in a node
func (t *BTree) minItems() int {
	return t.degree - 1
}
