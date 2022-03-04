package softheap

import (
	"fmt"
	"math"
	"strings"
)

type softHeapNode[T any] struct {
	currentKey int
	elements   []softHeapElement[T]
	heap       *SoftHeap[T]
	left       *softHeapNode[T]
	right      *softHeapNode[T]
	// the rank satisfies: `left.rank == right.rank == rank - 1`
	rank int
	// determines the approximate maximum length of `elements`:
	//      size/2 <= len(elements) <= 3*size
	size int
}

type softHeapElement[T any] struct {
	key   int
	value T
}

func newNode[T any](heap *SoftHeap[T], left, right *softHeapNode[T]) softHeapNode[T] {
	rank := 0
	size := 1
	if left != nil {
		rank = left.rank + 1
		if rank > r {
			size = int(math.Ceil(SizeFactor * float64(left.size)))
		}
	} else if right != nil {
		rank = right.rank + 1
		if rank > r {
			size = int(math.Ceil(SizeFactor * float64(right.size)))
		}
	}

	return softHeapNode[T]{
		heap:  heap,
		rank:  rank,
		left:  left,
		right: right,
		size:  size,
	}
}

func (n *softHeapNode[T]) pushElement(key int, value T) {
	element := softHeapElement[T]{key, value}
	n.elements = append(n.elements, element)
	if key > n.currentKey {
		n.currentKey = key
	}
}

func (n *softHeapNode[T]) popElement() (int, *T) {
	E := len(n.elements)
	if E == 0 {
		return -1, nil
	}
	e := n.elements[E-1]
	n.elements = n.elements[:E-1]
	return e.key, &e.value
}

func (n *softHeapNode[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *softHeapNode[T]) isEmpty() bool {
	return len(n.elements) == 0
}

// `sift` recursively moves elements from the `elements` list of child nodes
// the `elements` list of parent nodes. This allows us to remove "unneeded" leaf
// nodes to make the tree more compact
func (n *softHeapNode[T]) sift() {
	for len(n.elements) < n.size && !n.isLeaf() {
		if n.left == nil {
			n.left = n.right
			n.right = nil
		} else if n.right != nil && n.left.currentKey > n.right.currentKey {
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

// `siftIfNeeded` calls `sift` in case the length of the root
// list of elements is small. This allows us to move elements
// from the interior nodes to the root node, making it easy to retrieve.
// Returns true in case `sift` was actually needed.
func (n *softHeapNode[T]) siftIfNeeded() bool {
	if len(n.elements) <= n.size/2 {
		n.sift()
		return true
	}
	return false
}

// `combine` joins two trees by making them have a common parent.
// The tree returned will have `n` as its left child and `m` as its right child.
func (n *softHeapNode[T]) combine(m *softHeapNode[T]) *softHeapNode[T] {
	root := newNode(n.heap, n, m)
	root.sift()
	return &root
}

func (n *softHeapNode[T]) print(indent int) {
	fmt.Printf("%s[rank:%d; key: %d; size: %d; elts: %v]\n",
		strings.Repeat(" ", indent), n.rank, n.currentKey, n.size, n.elements)
	if n.left != nil {
		n.left.print(indent + 2)
	}
	if n.right != nil {
		n.right.print(indent + 2)
	}
}
