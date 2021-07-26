/* Shuffle elements from center to outwards in array */
package main

import (
	"fmt"
)

// Time: O(n)
func shuffle1(a []int, l, r int) {
	// return if 0 or 1 elements
	if l >= r {
		return
	}
	shuffle1(a, l+1, r-1)
	// don't swap for outer bounds
	if (r - l + 1) == len(a) {
		return
	}
	// observe temp for how merge occurs in each iteration from inside to outside
	for i := l; (i + 1) <= r; i = i + 2 {
		a[i], a[i+1] = a[i+1], a[i]
	}
	fmt.Println("temp: ", a)
}

/*   << Not Working >>  for few inputs eg: for [1 3 5 2 4 6] gives [1 2 3 5 6 4] */
// Time: O(logN)
func shuffle2(a []int, l, r int) {
	if l >= r {
		return
	}
	if r-l == 1 {
		return
	}
	mid := (l + r) / 2
	lmid := (l + mid) / 2
	for i, j := lmid+1, mid+1; i <= mid; i, j = i+1, j+1 {
		a[i], a[j] = a[j], a[i]
	}
	shuffle2(a, l, mid)
	shuffle2(a, mid+1, r)
}

func main() {
	a := []int{1, 3, 5, 2, 4, 6}
	fmt.Println("\nshuffle1 input: ", a)
	shuffle1(a, 0, len(a)-1)
	fmt.Println("shuffle1 output: ", a)
	a = []int{1, 3, 5, 7, 2, 4, 6, 8}
	fmt.Println("\nshuffle1 input: ", a)
	shuffle1(a, 0, len(a)-1)
	fmt.Println("shuffle1 output: ", a)

	/*
		a = []int{1, 3, 5, 2, 4, 6}
		fmt.Println("\n\nshuffle2 input: ", a)
		shuffle2(a, 0, len(a)-1)
		fmt.Println("shuffle2 output: ", a)
		a = []int{1, 3, 5, 7, 2, 4, 6, 8}
		fmt.Println("\nshuffle2 input: ", a)
		shuffle2(a, 0, len(a)-1)
		fmt.Println("shuffle2 output: ", a)
	*/
}
