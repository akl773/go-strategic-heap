package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		heap     *Heap[int]
		input    []int
		expected []int
	}{
		{
			name:     "Test MinHeap",
			heap:     NewMinHeap[int](),
			input:    []int{5, 2, 6, 1, 3, 4},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Test MaxHeap",
			heap:     NewMaxHeap[int](),
			input:    []int{1, 3, 5, 2, 6, 4},
			expected: []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Test Len initially and after each Push
			if gotLen := tc.heap.Len(); gotLen != 0 {
				t.Errorf("heap.Len() = %v, want 0", gotLen)
			}

			for i, v := range tc.input {
				tc.heap.Push(v)

				if gotLen := tc.heap.Len(); gotLen != i+1 {
					t.Errorf("heap.Len() = %v, want %v", gotLen, i+1)
				}
			}

			// Test Pop
			for _, v := range tc.expected {
				if got := tc.heap.Pop(); got != v {
					t.Errorf("heap.Pop() = %v, want %v", got, v)
				}
			}

			// Test Len after all Pops
			if gotLen := tc.heap.Len(); gotLen != 0 {
				t.Errorf("heap.Len() = %v, want 0", gotLen)
			}
		})
	}
}
