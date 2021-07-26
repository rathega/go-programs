/* GeeksForGeeks qn:
   Given an array A of integers. Find three numbers such that sum of two elements equals the third element and return the triplet in a container result, if no such triplet is found return the container as empty.
*/
package main

import (
	"fmt"
)

// Similar to tripletSum
// Time: O(n*n) Space: O(n)
// Keep one element(a) as fixed and check for other two elements.
// In set of three numbers that satisfy a, b, c there can be two possibility a = b+c or b = a+c (so checkNegative is done)
func tripletFamily(a []int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		m := make(map[int]struct{})
		for j := i + 1; j < n; j++ {
			ele := a[i] - a[j]
			// this checkNegative is the addition to tripletSum
			if ele < 0 {
				ele = -1 * ele
			}
			if _, ok := m[ele]; ok {
				fmt.Printf("\n%d %d %d", a[i], a[j], ele)
			} else {
				m[a[j]] = struct{}{}
			}
		}
	}
}

// Can also be done in Time: O(n*n) Space: 1 similar to func tripletSum3()

func main() {
	a := []int{1, 2, 3, 4, 5}
	tripletFamily(a)
}
