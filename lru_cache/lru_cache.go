package main

import (
	"fmt"
)

var queueMaxSize = 4	// the cache size
var hashMaxSize = 10	// range of pages referred here 0-9

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type Queue struct {
	Front *Node
	Rear *Node
	Capacity int
	CurSize int
}

type Hash struct {
	PageMap map[int]*Node
	Capacity int
}

// Queue functions
func (q *Queue) EnQueue(qEle *Node) {
	firstEle := q.Front
	if firstEle == nil {
		q.Front = qEle
		q.Rear = qEle
		return
	} else {
		firstEle.Prev = qEle
		qEle.Next = firstEle
		q.Front = qEle
	}
}

func (q *Queue) DeQueue() int {
	if q.Rear == nil {
		return -1
	}
	if q.Front == q.Rear {
		data := q.Front.Data
		q.Front = nil
		q.Rear = nil
		return data
	}
	data := q.Rear.Data
	newRear := q.Rear.Prev
	q.Rear.Prev.Next = nil
	q.Rear.Prev = nil
	q.Rear = newRear
	return data
}

func (q *Queue) AddToQueue(qEle *Node) int {
	lruPage := -1
	if q.CurSize == q.Capacity {
		// Max cache size reached remove lru page
		temp := q.DeQueue()
		lruPage = temp
	} else {
		q.CurSize++
	}
	q.EnQueue(qEle)
	return lruPage
}

func (q *Queue) BringToFront(qEle *Node) {
	if q.Front == q.Rear {
		return
	}
	if q.Front == qEle {
		return
	}
	if q.Rear == qEle {
		q.Rear = qEle.Prev
	}
	if qEle.Prev != nil {
		qEle.Prev.Next = qEle.Next
	}
	if qEle.Next != nil {
		qEle.Next.Prev = qEle.Prev
	}
	qEle.Next = q.Front
	q.Front.Prev = qEle
	q.Front = qEle
}

func (q *Queue) PrintQueue() {
	fmt.Println("Content of Queue:")
	qEle := q.Front
	for ;qEle != nil; {
		fmt.Println(qEle.Data)
		qEle = qEle.Next
	}
}

// Hash functions
func (h *Hash) ReferPage(q *Queue, pageNo int) (){
	fmt.Println("\nRefer Page: ", pageNo)
	if pageNo <  0 || pageNo >= hashMaxSize {
		fmt.Println("Invalid page")
		return
	}
	if _, ok := h.PageMap[pageNo]; !ok {
		// Page not in cache
		qEle := &Node{Data: pageNo, Prev: nil, Next: nil}
		lruPage := q.AddToQueue(qEle)
		if lruPage != -1 {
			delete(h.PageMap, lruPage)
		}
		h.PageMap[pageNo] = qEle
	} else {
		// Page presnt in cache bring to front
		q.BringToFront(h.PageMap[pageNo])
	}
	q.PrintQueue()
}

// Init LRU structs
func initLRU() (*Queue, *Hash){
	q := &Queue{Front:nil, Rear:nil, Capacity: queueMaxSize}
	h := &Hash{PageMap:make(map[int]*Node), Capacity: hashMaxSize}
	return q, h
}

// Main
func main() {
	queue, hash := initLRU()
	hash.ReferPage(queue, 1)
	hash.ReferPage(queue, 2)
	hash.ReferPage(queue, 3)
	hash.ReferPage(queue, 1)
	hash.ReferPage(queue, 4)
	hash.ReferPage(queue, 5)
}
