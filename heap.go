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

type Heap struct {
	data     []int
	strategy Strategy
}

func NewMinHeap() *Heap {
	return &Heap{data: nil, strategy: MinHeapStrategy{}}
}

func NewMaxHeap() *Heap {
	return &Heap{data: nil, strategy: MaxHeapStrategy{}}
}
