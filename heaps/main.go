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

func (h *MaxHeap) extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Can not extract from an empty array")
		return -1
	}

	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	childToCompare := 0

	leftChildIndex := getLeftChildIndex(index)
	rightChildIndex := getRightChildIndex(index)

	for leftChildIndex <= lastIndex {
		if (leftChildIndex == lastIndex) || (h.array[leftChildIndex] > h.array[rightChildIndex]) {
			childToCompare = leftChildIndex
		} else {
			childToCompare = rightChildIndex
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			leftChildIndex = getLeftChildIndex(index)
			rightChildIndex = getRightChildIndex(index)
		} else {
			return
		}

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

	buildHeap := []int{10, 20, 30, 40, 5, 8, 0, 13, 9, 7, 12, 46}

	for _, v := range buildHeap {
		m.insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 8; i++ {
		m.extract()
		fmt.Println(m)
	}
}
