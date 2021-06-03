package gobtree

import "sync"

const (
	defaultFreeListSize = 32
)

type Item interface {
	Less(other Item) bool
}

type node struct {
	items    []Item
	children []*node
	cow      *copyOnWriteContext
}

type FreeList struct {
	mu       sync.Mutex
	freelist []*node
}

func newFreeList(size int) *FreeList {
	return &FreeList{freelist: make([]*node, 0, size)}
}

func (f *FreeList) newNode() (n *node) {
	f.mu.Lock()
	defer f.mu.Unlock()
	last := len(f.freelist) - 1
	if last < 0 {
		return new(node)
	}
	n = f.freelist[last]
	f.freelist[last] = nil
	f.freelist = f.freelist[:last]
	return
}

type copyOnWriteContext struct {
	freelist *FreeList
}

func (c *copyOnWriteContext) newNode() (n *node) {
	n = c.freelist.newNode()
	n.cow = c
	return
}

type BTree struct {
	degree int
	size   int // number of elements in a whole tree
	root   *node
	cow    *copyOnWriteContext
}

func New(degree int) *BTree {
	return &BTree{
		degree: degree,
		cow: &copyOnWriteContext{
			freelist: newFreeList(defaultFreeListSize),
		},
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

func (t *BTree) Insert(item Item) Item {
	if item == nil {
		panic("nil cannot be added to a B-Tree")
	}
	if t.root == nil {
		t.root = t.cow.newNode()
		t.root.items = append(t.root.items, item)
		t.size++
		return nil
	}
	return item
}
