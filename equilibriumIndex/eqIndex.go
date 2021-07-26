/* GeeksForGeeks qn:
   Equilibrium index of an array is an index such that the sum of elements at lower indexes is equal to the sum of elements at higher indexes.
*/
package main

import "fmt"

// Time: O(n) -> Works only for +ve numbers
func eqIndexNotWorking(a []int) int {
	fmt.Println("\nInput: ", a)
	if len(a) == 0 {
		return -1
	}
	// init two pointers from left and right
	l := 0
	r := len(a) - 1
	// lsum, rsum contains sum till l, r elements on both sides
	lsum := a[l]
	rsum := a[r]
	for l < r {
		if (l+2) == r && lsum == rsum {
			// if found index in the middle of l and r, both sides sum is equal return middle point as eq point
			return l + 1
		} else if lsum < rsum {
			// if total more on right side add elements on left side
			l = l + 1
			lsum = lsum + a[l]
		} else {
			// if total more on left side add elements on right side
			r = r - 1
			rsum = rsum + a[r]
		}
	}
	// return -1 if not found point
	return -1
}

// Time: O(n) Works for both +ve and -ve numbers
func eqIndex(a []int) int {
	fmt.Println("\nInput: ", a)
	sum := 0
	// find the total sum in first pass
	for i := 0; i < len(a); i++ {
		sum = sum + a[i]
	}
	// lsum is calculated before comparison with rsum
	lsum := 0
	for i := 0; i < len(a); i++ {
		// current element is subtracted to get the rsum value
		sum = sum - a[i]
		// check if lsum and rsum are equal and return eq point index
		if lsum == sum {
			return i
		}
		// compute lsum including current element for next pass
		lsum = lsum + a[i]
	}
	// return -1 if not found point
	return -1
}

func main() {
	a := []int{1, 2, 0, 3}
	fmt.Printf("EqIndex: %d\n", eqIndex(a))
	a = []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Printf("EqIndex: %d\n", eqIndex(a))
}
