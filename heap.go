package softheap

import (
	"fmt"
	"math"
)

const ErrorParameter float64 = 1e-3
const SizeFactor float64 = 1.5

var r = int(math.Ceil(math.Log2(1/ErrorParameter))) + 2

type SoftHeap[T any] struct {
	// the heap keeps a doubly linked list
	// of binary trees.
	// the list is sorted by the rank of each tree
	treeListHead *softHeapTree[T]
	rank         int
}

func New[T any](initialKey int, initialValue T) SoftHeap[T] {
	return NewWithErrorParam(initialKey, initialValue)
}

func NewWithErrorParam[T any](initialKey int, initialValue T) SoftHeap[T] {
	heap := SoftHeap[T]{}
	node := newNode(&heap, nil, nil)
	node.pushElement(initialKey, initialValue)
	treeListHead := newTree(&heap, nil, nil, &node)
	heap.treeListHead = &treeListHead
	heap.rank = treeListHead.rank()
	return heap
}

func (h *SoftHeap[T]) Insert(key int, value T) {
	e := New(key, value)
	h.Meld(&e)
}

// `Meld` joins two soft heaps. The new larger heap is stored
// in `h`, while `i` is discarded.
func (h *SoftHeap[T]) Meld(i *SoftHeap[T]) {
	a := h // a references the heap with the smaller rank
	b := i // b references the heap with the larger rank
	if h.rank > i.rank {
		a = i
		b = h
	}

	at := a.treeListHead
	bt := b.treeListHead
	if at == nil && bt == nil {
		return
	}
	min := func() (m *softHeapTree[T]) {
		if at != nil && bt != nil {
			if at.rank() <= bt.rank() {
				m = at
				at = at.next
			} else {
				m = bt
				bt = bt.next
			}
		} else if at != nil {
			m = at
			at = at.next
		} else {
			m = bt
			bt = bt.next
		}
		m.heap = h
		return m
	}

	newListHead := min()
	newListTail := newListHead
	t := newListHead
	// Merge the ordered list of `a` with the ordereed list of `b` (the order
	// of the lists is by the rank of the root element).
	for at != nil || bt != nil {
		t.next = min()
		t.next.prev = t
		newListTail = t
		t = t.next
	}

	t = newListHead
	// Combine consecutive trees with the same rank, unless we have
	// three consecutive trees with the same rank, in which case
	// only the last two are combined.
	for t != nil && t.next != nil {
		if t.rank() == t.next.rank() &&
			(t.next.next == nil || t.rank() != t.next.next.rank()) {
			t.combine()
		}
		t = t.next
	}

	h.treeListHead = newListHead
	h.rank = b.rank
	newListTail.updateSuffixMin()
}

// `ExtractMin` returns the element with the smallest
// current key, and deletes it from the heap.
func (h *SoftHeap[T]) ExtractMin() *T {
	if h.treeListHead == nil {
		return nil
	}
	s := h.treeListHead.suffmin
	_, e := s.extractMin()
	if s.isEmpty() {
		prev := s.prev
		// remove s from the linked list
		if s.prev != nil {
			s.prev.next = s.next
		} else {
			h.treeListHead = s.next
		}
		if s.next != nil {
			s.next.prev = s.prev
		}
		if prev != nil {
			// `s` was removed and it was the `suffmin` value
			// of the elements preceding it. Therefore we
			// need to update the `suffmin` of those elements.
			prev.updateSuffixMin()
		}
	}
	return e
}

func (h *SoftHeap[T]) Delete(key int) {
}

func (h *SoftHeap[T]) Print() {
	t := h.treeListHead
	fmt.Println("--------[ Heap ]--------")
	for t != nil {
		fmt.Printf("Tree: rank=%d suffmin.rank=%d\n",
			t.rank(), t.suffmin.rank())
		t.print()
		t = t.next
	}
}
