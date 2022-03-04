package softheap

type softHeapTree[T any] struct {
	prev *softHeapTree[T]
	next *softHeapTree[T]
	heap *SoftHeap[T]
	root *softHeapNode[T]
	// `suffmin` points to the tree in front of this tree in
	// the linked-list whose root has the smallest value
	// of `currentKey`. For example if this tree is T1 and we have
	// T1 -> T2 -> T3 -> T4 -> T5
	// and the value of `T3.root.currentKey` is smaller
	// than `Ti.root.currentKey` for each 2 < i < 5,
	// then `T1.suffmin = T3`.
	suffmin *softHeapTree[T]
}

func newTree[T any](heap *SoftHeap[T], prev, next *softHeapTree[T],
	root *softHeapNode[T]) softHeapTree[T] {
	t := softHeapTree[T]{prev: prev, next: next, heap: heap, root: root}
	t.suffmin = &t
	return t
}

func (t *softHeapTree[T]) rank() int {
	return t.root.rank
}

// `combine` joins a tree with the tree in front of it in the
// linked list of trees. The two consecutive elements must have
// the same rank (i.e. their root elements must have the same rank)
func (t *softHeapTree[T]) combine() {
	// join the corresponding trees
	t.root = t.root.combine(t.next.root)
	// join the two consecutive linked-list nodes
	t.next = t.next.next
	if t.next != nil {
		t.next.prev = t
	}
}

// `updateSuffixMin` updates the value of `t.suffmin`.
// the function proceeds recursively on the list from tail to head,
// updating each node. On processing `t`, the procedure assumes that
// the values of suffmin have been updated for all nodes in front of `t`.
// After updating `t`, the function is recursively called on `t.prev`
func (t *softHeapTree[T]) updateSuffixMin() {
	if t.next == nil || t.root.currentKey <= t.next.root.currentKey {
		t.suffmin = t
	} else {
		t.suffmin = t.next.suffmin
	}
	if t.prev != nil {
		t.prev.updateSuffixMin()
	}
}

// `extractMin` obtains an element in the tree with the
// smallest key. This element will be in the list of
// elements of the root node.
func (t *softHeapTree[T]) extractMin() (int, *T) {
	k, v := t.root.popElement()
	// move new elements to the root node if needed
	if t.root.siftIfNeeded() {
		// the current key of the root might have changed because
		// of the sift so we may need need to update the `suffmin` value
		t.updateSuffixMin()
	}
	return k, v
}

func (t *softHeapTree[T]) isEmpty() bool {
	return t.root.isEmpty()
}

func (t *softHeapTree[T]) print() {
	t.root.print(2)
}
