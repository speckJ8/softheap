package softheap

type SoftHeapNode[T any] struct {
	OriginalKey int
	CurrentKey  int
	Value       T
}

func (n *SoftHeapNode[T]) Corrupted() bool {
	return n.CurrentKey > n.OriginalKey
}
