package heap

type Strategy[T INumbers] interface {
	ShouldSwap(parent, child T) bool
}

type MinHeapStrategy[T INumbers] struct{}

func (MinHeapStrategy[T]) ShouldSwap(parent, child T) bool {
	return parent > child
}

type MaxHeapStrategy[T INumbers] struct{}

func (MaxHeapStrategy[T]) ShouldSwap(parent, child T) bool {
	return parent < child
}

type INumbers interface {
	int | float64
}

type Heap[T INumbers] struct {
	data     []T
	strategy Strategy[T]
}

func NewMinHeap[T INumbers]() *Heap[T] {
	return &Heap[T]{data: nil, strategy: MinHeapStrategy[T]{}}
}

func NewMaxHeap[T INumbers]() *Heap[T] {
	return &Heap[T]{data: nil, strategy: MaxHeapStrategy[T]{}}
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up(h.Len() - 1)
}

func (h *Heap[T]) Pop() T {
	minVal := h.data[0]
	h.swap(0, h.Len()-1)
	h.data = h.data[:h.Len()-1]
	h.down(0)
	return minVal
}

func (h *Heap[T]) isInBound(index int) bool {
	return index >= 0 && index < h.Len()
}

func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) up(i int) {
	for i > 0 && h.strategy.ShouldSwap(h.data[(i-1)/2], h.data[i]) {
		h.swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *Heap[T]) down(i int) {
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
