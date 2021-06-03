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
	length int
	root   *node
}
