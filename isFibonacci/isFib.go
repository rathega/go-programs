package main

import (
	"fmt"
	"math"
)

// function to return if a number is square number
func isSquare(a int) bool {
	sqrt := int(math.Sqrt(float64(a)))
	if sqrt*sqrt == a {
		return true
	}
	return false
}

// if 5*a*a-4 (or) 5*a*a+4 is a square number then a is a fibonacci number
func isFib(a int) {
	temp := 5 * a * a
	if isSquare(temp-4) || isSquare(temp+4) {
		fmt.Printf("%d is Fib\n", a)
	} else {
		fmt.Printf("%d is Not Fib\n", a)
	}
}

func main() {
	a := []int{4, 2, 8, 5, 20, 1, 40, 13, 23}
	for _, ele := range a {
		isFib(ele)
	}
}
