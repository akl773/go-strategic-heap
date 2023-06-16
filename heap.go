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

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) isInBound(index int) bool {
	return index >= 0 && index < h.Len()
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) up(i int) {
	for i > 0 && h.strategy.ShouldSwap(h.data[(i-1)/2], h.data[i]) {
		h.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *Heap) down(i int) {
	for {
		left, right := 2*i+1, 2*i+2
		minIndex := i

		if h.isInBound(left) && h.strategy.ShouldSwap(h.data[minIndex], h.data[left]) {
			minIndex = left
		}
		if h.isInBound(right) && h.strategy.ShouldSwap(h.data[minIndex], h.data[right]) {
			minIndex = right
		}
		if minIndex != i {
			h.swap(i, minIndex)
			i = minIndex
		} else {
			break
		}
	}
}
