/* GeeksForGeeks qn:
   A Bitonic Sequence is a sequence of numbers which is first strictly increasing then after a point strictly decreasing.
   A Bitonic Point is a point in bitonic sequence before which elements are strictly increasing and after which elements are strictly decreasing. A Bitonic point doesn’t exist if array is only decreasing or only increasing.
*/
package main

import (
	"fmt"
)

func bitonic(a []int, l, r int) int {
	// if less than three elements in array won't enter for loop and directly returns -1
	for l <= r {
		mid := (l + r) / 2
		if a[mid-1] < a[mid] && a[mid] > a[mid+1] {
			return a[mid]
		} else if a[mid-1] > a[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func bitonicRec(a []int, l, r int) int {
	// ending condition for recursion
	if l > r {
		return -1
	}
	mid := int((l + r) / 2)
	if a[mid-1] < a[mid] && a[mid] > a[mid+1] {
		return a[mid]
	} else if a[mid-1] > a[mid] {
		return bitonicRec(a, l, mid-1)
	} else {
		return bitonicRec(a, mid+1, r)
	}
}

// since bitonic point is only possible with atleast three elements pass 1 to n-2 index as boundaries
func main() {
	a := []int{1, 15, 25, 45, 42, 21, 17, 12, 11}
	fmt.Println("\nInput: ", a)
	fmt.Printf("Output: %d\n", bitonic(a, 1, len(a)-2))
	fmt.Printf("OutputRec: %d\n", bitonicRec(a, 1, len(a)-2))
	a = []int{1, 45, 47, 50, 5}
	fmt.Println("\nInput: ", a)
	fmt.Printf("Output: %d\n", bitonic(a, 1, len(a)-2))
	fmt.Printf("OutputRec: %d\n", bitonicRec(a, 1, len(a)-2))
	a = []int{1, 2, 3, 4, 5}
	fmt.Println("\nInput: ", a)
	fmt.Printf("Output: %d\n", bitonic(a, 1, len(a)-2))
	fmt.Printf("OutputRec: %d\n", bitonicRec(a, 1, len(a)-2))
	a = []int{5, 4, 3, 2, 1}
	fmt.Println("\nInput: ", a)
	fmt.Printf("Output: %d\n", bitonic(a, 1, len(a)-2))
	fmt.Printf("OutputRec: %d\n", bitonicRec(a, 1, len(a)-2))
}
