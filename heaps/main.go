package main

import "fmt"

// Struct to hold the heap data structure
type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 40}

	for _, v := range buildHeap {
		m.insert(v)
	}
	fmt.Println(m)
}
