package main

import (
	"fmt"
)

// Queue structure
type Queue struct {
	front int   // front of queue
	rear  int   // rear of queue
	size  int   // size of queue
	arr   []int // array to store elements
}

// Initialise queue for given size
func NewQueue(size int) *Queue {
	q := &Queue{}
	q.front = -1
	q.rear = -1
	q.size = size
	q.arr = make([]int, size)
	return q
}

// Enqueue to add elements at rear end
func (q *Queue) Enqueue(ele int) {
	fmt.Printf("Enqueue %d: ", ele)
	// check for overflow
	if q.rear == q.size-1 {
		fmt.Println("Queue is full!")
		return
	}
	// increment rear and add
	q.rear++
	q.arr[q.rear] = ele
	// if adding first element set front
	if q.front == -1 {
		q.front = 0
	}
	q.Display()
}

// Dequeue to delete elements from front end
func (q *Queue) Dequeue() {
	fmt.Printf("Dequeue: ")
	// check for empty queue
	if q.front == -1 {
		fmt.Println("Queue is empty!")
		return
	}
	// first element at front is removed
	fmt.Printf("%d ", q.arr[q.front])
	// if only element unset front, rear
	if q.front == q.rear {
		q.arr[q.front] = 0
		q.front = -1
		q.rear = -1
		q.Display()
		return
	}
	// move all elements one place left to fill previous removed space
	// Note: this helps to utilise full queue size instead of waiting for
	//       all elements to be removed to retrieve empty space
	for i := q.front; i < q.rear; i++ {
		q.arr[i] = q.arr[i+1]
	}
	// decrement rear pointer after moving elements to left
	q.rear--
	q.Display()
}

// Display to print queue elements from front to rear
func (q *Queue) Display() {
	fmt.Printf("front -> ")
	if q.front == -1 {
		fmt.Println("<- rear")
		return
	}
	for i := q.front; i <= q.rear; i++ {
		fmt.Printf("%d ", q.arr[i])
	}
	fmt.Println("<- rear")
}

func main() {
	q := NewQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
}
