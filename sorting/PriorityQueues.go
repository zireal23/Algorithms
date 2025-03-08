package sorting

import "fmt"

// PriorityQueue implements a max-heap based priority queue
type PriorityQueue struct {
	items []int
	size  int
}

// NewPriorityQueue creates a new empty priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		items: make([]int, 1), // Index 0 unused for simpler parent/child calculations
		size:  0,
	}
}

// NewPriorityQueueFromArray creates a priority queue from an existing array
func NewPriorityQueueFromArray(items []int) *PriorityQueue {
	pq := NewPriorityQueue()
	for _, item := range items {
		pq.Insert(item)
	}
	return pq
}

// Insert adds a new item to the priority queue
func (pq *PriorityQueue) Insert(item int) {
	// Resize if needed
	if pq.size >= len(pq.items)-1 {
		newSize := max(2*len(pq.items), 2)
		newItems := make([]int, newSize)
		copy(newItems, pq.items)
		pq.items = newItems
	}

	pq.size++
	pq.items[pq.size] = item
	pq.swimUp()
}

// RemoveMax removes and returns the maximum element from the queue
// Returns the max element and an error if the queue is empty
func (pq *PriorityQueue) RemoveMax() (int, error) {
	if pq.size == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	max := pq.items[1]
	pq.items[1] = pq.items[pq.size]
	pq.size--
	pq.sinkDown()
	return max, nil
}

// Size returns the number of elements in the queue
func (pq *PriorityQueue) Size() int {
	return pq.size
}

// HeapSort sorts the elements in the queue using heap sort
func (pq *PriorityQueue) HeapSort() []int {
	sorted := make([]int, pq.size)
	for i := pq.size; i > 0; i-- {
		max, err := pq.RemoveMax()
		if err != nil {
			break
		}
		sorted[i-1] = max
	}
	return sorted
}

// swimUp maintains the heap property by moving the last element up
func (pq *PriorityQueue) swimUp() {
	k := pq.size
	for k > 1 && pq.items[k/2] < pq.items[k] {
		pq.items[k/2], pq.items[k] = pq.items[k], pq.items[k/2]
		k = k / 2
	}
}

// sinkDown maintains the heap property by moving the root element down
func (pq *PriorityQueue) sinkDown() {
	k := 1
	for 2*k <= pq.size {
		j := 2*k
		if j < pq.size && pq.items[j] < pq.items[j+1] {
			j++
		}
		if pq.items[k] >= pq.items[j] {
			break
		}
		pq.items[k], pq.items[j] = pq.items[j], pq.items[k]
		k = j
	}
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
