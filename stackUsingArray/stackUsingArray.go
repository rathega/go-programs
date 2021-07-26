package main

import (
	"fmt"
)

// Stack structure
type Stack struct {
	tos  int   // top of stack
	size int   // size of stack
	arr  []int // array to store elements
}

// Initialise stack for given size
func NewStack(size int) *Stack {
	s := &Stack{}
	s.tos = -1
	s.size = size
	s.arr = make([]int, size)
	return s
}

// Push elements to tos
func (s *Stack) Push(ele int) {
	fmt.Printf("Push %d: ", ele)
	// check for overflow
	if s.tos == s.size-1 {
		fmt.Println("Stack is full!")
		return
	}
	// increment tos and add element
	s.tos++
	s.arr[s.tos] = ele
	s.PrintStack()
}

// Pop elements from tos
func (s *Stack) Pop() {
	fmt.Printf("Pop: ")
	// check if stack is empty
	if s.tos == -1 {
		fmt.Println("Stack is empty!")
		return
	}
	// take element at tos and decrement tos
	fmt.Printf("%d ", s.arr[s.tos])
	s.tos--
	s.PrintStack()
}

// PrintStack to see contents of stack
func (s *Stack) PrintStack() {
	fmt.Printf("TOS: %d Stack: ", s.tos)
	for i := s.tos; i >= 0; i-- {
		fmt.Printf("%d ", s.arr[i])
	}
	fmt.Println()
}

func main() {
	s := NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
}
