/* MinHeap implementation using array*/
package main

import (
	"fmt"
)

type MinHeap struct {
	arr  []int // slice to store heap elements
	size int   // size denotes heap capacity
	last int   // last denotes current last valid heap index
}

func (h *MinHeap) Init(size int) {
	fmt.Printf("Init MinHeap with size: %d\n", size)
	h.size = size
	h.arr = make([]int, size)
	h.last = -1
	h.PrintHeap()
}

func (h *MinHeap) Insert(ele int) {
	fmt.Printf("Insert element %d: ", ele)

	// check if heap is full
	if h.last >= h.size-1 {
		fmt.Println("Heap is full")
		return
	}

	// insert element at end
	h.last++
	h.arr[h.last] = ele

	// heapify from bottom to top
	for i := h.last; (i-1)/2 >= 0; {
		parent := (i - 1) / 2
		if h.arr[parent] > h.arr[i] {
			temp := h.arr[parent]
			h.arr[parent] = h.arr[i]
			h.arr[i] = temp
			i = parent
		} else {
			break
		}
	}

	h.PrintHeap()
}
func (h *MinHeap) Delete() {
	fmt.Printf("Delete Min: ")

	// check if heap is empty
	if h.last <= -1 {
		fmt.Println("Heap is empty")
		return
	}

	// take the first element (min element)
	ele := h.arr[0]
	fmt.Println(ele)

	// make the last element as first element
	h.arr[0] = h.arr[h.last]
	h.last--

	// heapify from top to bottom
	for i := 0; 2*i <= h.last; {
		lchild := 2*i + 1
		rchild := 2*i + 2
		if lchild <= h.last && h.arr[lchild] < h.arr[i] {
			temp := h.arr[i]
			h.arr[i] = h.arr[lchild]
			h.arr[lchild] = temp
			i = lchild
			continue
		} else if rchild <= h.last && h.arr[rchild] < h.arr[i] {
			temp := h.arr[i]
			h.arr[i] = h.arr[rchild]
			h.arr[rchild] = temp
			i = rchild
			continue
		}
		break
	}

	h.PrintHeap()
}

func (h *MinHeap) PrintHeap() {
	fmt.Printf("[ ")
	for i := 0; i <= h.last; i++ {
		fmt.Printf("%d ", h.arr[i])
	}
	fmt.Println("]")
}

func main() {
	h := &MinHeap{}
	h.Init(5)
	h.Insert(5)
	h.Insert(4)
	h.Insert(3)
	h.Insert(2)
	h.Insert(1)
	h.Insert(0)
	h.Delete()
	h.Delete()
	h.Delete()
	h.Delete()
	h.Delete()
	h.Delete()
}
