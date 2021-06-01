/* Rotating array clockwise and anti-clockwise */
package main

import (
	"fmt"
)

func clockRotate(inp []int, d int) {
	fmt.Printf("Clockwise rotate %v by %d\n", inp, d)

	// copy input to local slice - needed to avoid overwrite in inp array
	a := make([]int, len(inp))
	copy(a, inp)

	// temp array to store replacing elements
	b := []int{}

	// to avoid redundant rotate
	d = d % len(a)

	// a:1 2 3 4 5 d:1
	// b:
	for i := len(a) - d; i < len(a); i++ {
		b = append(b, a[i])
	}
	// a:1 2 3 4 5 d:1
	// b:5

	// a:1 2 3 4 5 d:1
	// b:5
	for i := len(a) - 1; (i - d) >= 0; i-- {
		a[i] = a[i-d]
	}
	// a:1 1 2 3 4
	// b:5

	// a:1 1 2 3 4
	// b:5
	for i := 0; i < d; i++ {
		a[i] = b[i]
	}
	// a:5 1 2 3 4
	fmt.Println(a)
}

func antiClockRotate(inp []int, d int) {
	fmt.Printf("Anti-Clockwise rotate %v by %d\n", inp, d)

	// copy input to local slice - needed to avoid overwrite in inp array
	a := make([]int, len(inp))
	copy(a, inp)

	// temp array to store replacing elements
	b := []int{}

	// to avoid redundant rotate
	d = d % len(a)

	// a:1 2 3 4 5 d:1
	// b:
	for i := 0; i < d; i++ {
		b = append(b, a[i])
	}
	// a:1 2 3 4 5 d:1
	// b:1

	// a:1 2 3 4 5 d:1
	// b:1
	for i := 0; (i + d) < len(a); i++ {
		a[i] = a[i+d]
	}
	// a:2 3 4 5 5
	// b:1

	// a:2 3 4 5 5
	// b:1
	for i := len(a) - d; i < len(a); i++ {
		a[i] = b[0]
		b = b[1:]
	}
	// a:2 3 4 5 1
	// b:
	fmt.Println(a)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	clockRotate(a, 0)
	clockRotate(a, 1)
	clockRotate(a, 2)
	clockRotate(a, 3)
	clockRotate(a, 4)
	clockRotate(a, 5)
	b := []int{1, 2, 3, 4, 5}
	antiClockRotate(b, 0)
	antiClockRotate(b, 1)
	antiClockRotate(b, 2)
	antiClockRotate(b, 3)
	antiClockRotate(b, 4)
	antiClockRotate(b, 5)
}
