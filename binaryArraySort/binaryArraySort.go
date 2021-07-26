/* GeeksForGeeks qn:

   Given a binary array A[] of size N. The task is to arrange the array in increasing order.

   Note: The binary array contains only 0  and 1.
*/

package main

import "fmt"

// this geeksforgeeks found algo fails for [0 1 0 1 1 0 0] gives output [0 0 0 1 1 1 0]
// -> this is because if first element is 0 then swapping the first element moves 0 towards right!
func binarySortNotWorking(inp []int, n int) {
	a := make([]int, n)
	copy(a, inp)
	fmt.Println(a)
	zeroIndex := 0
	for i := 1; i < n; i++ {
		if a[i] == 0 {
			a[zeroIndex], a[i] = a[i], a[zeroIndex]
			zeroIndex++
		}
	}
	fmt.Println(a)
}

// Binary Sort
func binarySort(inp []int, n int) {
	// []int copied to another array to avoid overwriting on input
	a := make([]int, n)
	copy(a, inp)
	fmt.Println("Input: ", a)

	// start two indexes one from first, another from last
	for l, r := 0, n-1; l <= r; l, r = l+1, r-1 {
		// move towards right until 1 is found
		for a[l] == 0 && l < r {
			l++
		}
		// move towards left until 0 is found
		for a[r] == 1 && l < r {
			r--
		}
		// swap 1 on left and 0 on right
		if l < r {
			a[l], a[r] = a[r], a[l]
		}
	}
	fmt.Println("Binary sort: ", a)
}

func main() {
	a := []int{0, 1, 0, 1, 1, 0, 0}
	binarySort(a, len(a))
}
