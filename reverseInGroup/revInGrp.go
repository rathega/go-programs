/* GeeksForGeeks qn:
   Given an array arr[] of positive integers of size N. Reverse every sub-array group of size K.
*/

package main

import "fmt"

func revInGrp(inp []int, k int) {
	// copy to temp array to avoid overwrite on original inp
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("Input: ", a)

	// consider input in groups of k
	for i := 0; i < len(a); i = i + k {
		// j is ending index of current group
		j := i + k - 1
		// this is needed if last group in input size is < k
		if j >= len(a) {
			j = len(a) - 1
		}
		// i, j are the bounds of current group so reverse in this for loop
		for tempi := i; tempi < j; tempi, j = tempi+1, j-1 {
			a[tempi], a[j] = a[j], a[tempi]
		}
	}
	fmt.Printf("Reverse in group of %d: %v\n", k, a)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	revInGrp(a, 1)
	revInGrp(a, 2)
	revInGrp(a, 3)
	revInGrp(a, 4)
	revInGrp(a, 5)
	revInGrp(a, 6)
}
