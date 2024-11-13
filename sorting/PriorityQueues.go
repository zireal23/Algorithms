package sorting

import "fmt"

// Writing an industrial strength priority queue implementation:
//     1. We can recieve a stream of integers that we build a heap of
//     2. We can recieve an array of integers to build a help of

type PriorityQueue struct {
	pq []int
	N  int
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		pq: make([]int, 1),
		N:  0,
	}
}

func NewPriorityQueueFromArray(intList []int) *PriorityQueue {
	pq := NewPriorityQueue()
	for _, item := range intList {
		AddItemToList(pq, item)
	}
	return pq
}

// Add a new item to the heap
//  1. We check if the existing capacity of the underlying array has the capacity
//     a. We resize it once the number of elements in the array is half the capacity of the array
//     b. We do this to amortize the cost of array resize that will inevtiably happen when our heap grows
//  2. Now we add the newly inserted element to the end of the heap
//  3. This disrupts the heap invariant and to restore that From here we make it "SWIM UP" to its proper place so that the heap invariant is maintained
//
// Add a new item to the heap
//  1. We check if the existing capacity of the underlying array has the capacity
//     a. We resize it once the number of elements in the array is half the capacity of the array
//     b. We do this to amortize the cost of array resize that will inevtiably happen when our heap grows
//  2. Now we add the newly inserted element to the end of the heap
//  3. This disrupts the heap invariant and to restore that From here we make it "SWIM UP" to its proper place so that the heap invariant is maintained
func AddItemToList(pq *PriorityQueue, item int) {
	if pq.N >= len(pq.pq)-1 {
		oldpq := pq.pq
		newSize := max(2*len(oldpq), 2) // Ensure at least size 2 on first resize
		pq.pq = make([]int, newSize)
		copy(pq.pq, oldpq)
	}
	// Add the new item to the end of the heap
	pq.N++
	pq.pq[pq.N] = item // Place the new item at the last position
	// Restore the heap property by swimming the new item up
	swimUp(pq)
}

func RemoveMaxElement(pq *PriorityQueue) int {
	maxElement := pq.pq[1]
	fmt.Printf("Max element is: %d\n", maxElement)
	Exch(pq.pq, 1, pq.N)
	pq.N--
	sinkDown(pq)
	return maxElement
}

func HeapSort(pq *PriorityQueue) []int {
	heapLen := pq.N
	sortedArray := make([]int, heapLen-1)
	for i := 0; i < heapLen-1; i++ {
		sortedArray[i] = RemoveMaxElement(pq)
	}
	return sortedArray
}

// We always check for the last element for the swimup since we enforce non immutable keys for our priority queue implementation
//  1. We start from our last element that is at the end of our complete binary tree(heap)
//  2. Now we check the parent element of our last element
//     a. If its bigger, then the max heap invariant holds and break out
//     b. If not, then we exchange both the elements and continue the search for enforcing the heap invariant
func swimUp(pq *PriorityQueue) {
	k := pq.N / 2
	lastIndex := pq.N
	for k > 0 {
		// fmt.Printf("Current element at index: %d and element is: %d\n", pq.N, pq.pq[lastIndex])
		// fmt.Printf("Current parent element at index: %d and element is: %d\n", k, pq.pq[k])
		currItem := pq.pq[lastIndex]
		parentItem := pq.pq[k]
		if parentItem > currItem {
			break
		}
		Exch(pq.pq, lastIndex, k)
		lastIndex = k
		k = k / 2
	}
}

// We use this function when we want to remove any element from the top of the priority queue
//  1. We take the first element of the queue, take the last element of the queue and put it at the top
//  2. Then we keep comparing the top element with its children:
//     a. If its greater than both the children then we break out of the loop
//     b. If its smaller than one or both the children then we exchange it with its greater children
//  3. We keep on looping until we reach the point where it is the greater than both its children
func sinkDown(pq *PriorityQueue) {
	currInd := 1
	for currInd*2 <= pq.N {
		targetChild := currInd * 2
		if targetChild < pq.N && pq.pq[targetChild] < pq.pq[targetChild+1] {
			targetChild++
		}
		if pq.pq[currInd] >= pq.pq[targetChild] {
			break
		}
		Exch(pq.pq, currInd, targetChild)
		currInd = targetChild
	}

}
