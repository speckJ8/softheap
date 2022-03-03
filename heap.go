package softheap

const DefaultErrorParameter = 1e-10

type SoftHeap[T any] struct {
	// error paremeter
	e float32
}

func New[T any]() SoftHeap[T] {
	return SoftHeap[T]{e: DefaultErrorParameter}
}

func NewWithErrorParam[T any](e float32) SoftHeap[T] {
	return SoftHeap[T]{e}
}

func (h *SoftHeap[T]) Insert(key int, value T) {
}

func (h *SoftHeap[T]) Delete(key int) {
}

// `Meld` joins two soft heaps, returning the new larger heap.
// The initial heaps are invalidated in the process.
func (h *SoftHeap[T]) Meld(e *SoftHeap[T]) SoftHeap[T] {
	return SoftHeap[T]{}
}

// `ExtractMin` returns the element with the smallest
// current key, and deletes it from the heap.
// func (h *SoftHeap[T]) ExtractMin() T {
// return nil
// }
