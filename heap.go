package heap

type Strategy interface {
	ShouldSwap(parent, child int) bool
}

type MinHeapStrategy struct{}

func (MinHeapStrategy) ShouldSwap(parent, child int) bool {
	return parent > child
}

type MaxHeapStrategy struct{}

func (MaxHeapStrategy) ShouldSwap(parent, child int) bool {
	return parent < child
}
