package softheap

import (
	"math"
)

const sizeFactor float64 = 1.5

type softHeapNode[T any] struct {
	currentKey int
	elements   []softHeapElement[T]
	heap       *SoftHeap[T]
	parent     *softHeapNode[T]
	left       *softHeapNode[T]
	right      *softHeapNode[T]
	// the rank satisfies: `parent.rank == rank + 1` and
	// `left.rank == right.rank == rank - 1`
	rank int
	// determines the maximum length of `elements`
	size int
}

type softHeapElement[T any] struct {
	key   int
	value T
}

func newNode[T any](heap *SoftHeap[T], parent, left, right *softHeapNode[T]) softHeapNode[T] {
	rank := 0
	size := 1
	if left != nil {
		rank = left.rank + 1
		size = int(math.Ceil(sizeFactor * float64(left.size)))
	} else if right != nil {
		rank = right.rank + 1
		size = int(math.Ceil(sizeFactor * float64(right.size)))
	}
	return softHeapNode[T]{
		heap:   heap,
		rank:   rank,
		parent: parent,
		left:   left,
		right:  right,
		size:   size,
	}
}

func (n *softHeapNode[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

// `sift` recursively moves elements from the `elements` list of child nodes
// the `elements` list of parent nodes. This allows us to remove "unneeded" leaf
// nodes to make the tree more compact
func (n *softHeapNode[T]) sift() {
	for len(n.elements) < n.size && !n.isLeaf() {
		if n.left == nil {
			n.left = n.right
		} else if n.left.currentKey > n.right.currentKey {
			tmp := n.left
			n.left = n.right
			n.right = tmp
		}
		// NOTE The reference "A simpler implementation and analysis of Chazelle’s Soft
		// Heaps" (Kaplan, Zwick) assumes that this is done in constant time by joining two
		// linked lists, but currently we are using arrays for the list of elements.
		// Therefore, the complexity measure of this implementation does not properly match
		// that of the paper.
		// TODO: change the list of elements to a linked list in the future...
		n.elements = append(n.elements, n.left.elements...)
		n.left.elements = nil
		n.currentKey = n.left.currentKey
		if !n.left.isLeaf() {
			n.left.sift()
		} else {
			n.left = nil
		}
	}
}

// `combine` joins two trees by making them have a common parent.
// The tree returned will have `n` as its left child and `m` as its right child.
func (n *softHeapNode[T]) combine(m *softHeapNode[T]) *softHeapNode[T] {
	root := newNode(n.heap, nil, n, m)
	root.left.parent = &root
	root.right.parent = &root
	root.sift()
	return &root
}
