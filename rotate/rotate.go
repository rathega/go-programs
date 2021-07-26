/* Rotating array clockwise and anti-clockwise */
package main

import (
	"fmt"
)

// Time: O(n) Space: O(d)
func clockRotate(inp []int, d int) {
	fmt.Printf("Clockwise rotate %v by %d:    ", inp, d)

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

// Time: O(n) Space: O(1) Juggling algorithm
func clockRotateJug(inp []int, d int) {
	fmt.Printf("Clockwise Jug rotate %v by %d:  ", inp, d)

	// copy input to local slice - needed to avoid overwrite in inp array
	a := make([]int, len(inp))
	copy(a, inp)

	g := gcd(len(a), d)
	for i := 0; i < g; i++ {
		temp := a[i]
		j := i
		k := i - d
		if k < 0 {
			k = k + len(a)
		}
		for k != i {
			a[j] = a[k]
			j = k
			k = k - d
			if k < 0 {
				k = k + len(a)
			}
		}
		a[j] = temp
	}
	fmt.Println(a)
}

// Time: O(n) Space: O(d)
func antiClockRotate(inp []int, d int) {
	fmt.Printf("Anti-Clockwise rotate %v by %d:    ", inp, d)

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

// Time: O(n) Space: O(1) Juggling algorithm
func antiClockRotateJug(inp []int, d int) {
	fmt.Printf("Anti-Clockwise Jug rotate %v by %d:  ", inp, d)

	// copy input to local slice - needed to avoid overwrite in inp array
	a := make([]int, len(inp))
	copy(a, inp)

	g := gcd(len(a), d)
	for i := 0; i < g; i++ {
		temp := a[i]
		j := i
		k := (i + d) % len(a)
		for ; k != i; j, k = k, (k+d)%len(a) {
			a[j] = a[k]
		}
		a[j] = temp
	}
	fmt.Println(a)
}

func gcd(a, b int) int {
	g := 1
	if a > b {
		g = b
	} else {
		g = a
	}
	for i := g; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	clockRotate(a, 0)
	clockRotate(a, 1)
	clockRotate(a, 2)
	clockRotate(a, 3)
	clockRotate(a, 4)
	clockRotate(a, 5)
	clockRotateJug(a, 0)
	clockRotateJug(a, 1)
	clockRotateJug(a, 2)
	clockRotateJug(a, 3)
	clockRotateJug(a, 4)
	clockRotateJug(a, 5)
	b := []int{1, 2, 3, 4, 5}
	antiClockRotate(b, 0)
	antiClockRotate(b, 1)
	antiClockRotate(b, 2)
	antiClockRotate(b, 3)
	antiClockRotate(b, 4)
	antiClockRotate(b, 5)
	antiClockRotateJug(b, 0)
	antiClockRotateJug(b, 1)
	antiClockRotateJug(b, 2)
	antiClockRotateJug(b, 3)
	antiClockRotateJug(b, 4)
	antiClockRotateJug(b, 5)
}
