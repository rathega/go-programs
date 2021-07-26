/* GeeksForGeeks qn:
   Given an array Arr[], write a program that segregates even and odd numbers. The program should put all even numbers first in sorted order, and then odd numbers in sorted order.
*/
package main

import (
	"fmt"
	"sort"
)

// Time: O(n) + O(nLogn) = O(nLogn) Space: O(1)
// Find an odd no. in first and even no. in last and swap
// Once odd even nos. are separated apply sort separately on two halves
func segregateEvenOdd(a []int) {
	fmt.Println("Input: ", a)

	// set two pointers l to first and r to last
	l := 0
	r := len(a) - 1
	for l < r {
		// loop till odd number is found on left
		for l < r && a[l]%2 == 0 {
			l = l + 1
		}
		// loop till even number is found on right
		for l < r && a[r]%2 != 0 {
			r = r - 1
		}
		// swap left and right nos.
		if l < r {
			a[l], a[r] = a[r], a[l]
			l = l + 1
			r = r - 1
		}
		// to find the dividing boundary of odd even nos.
		evenNo := 0
		// count the no. of even nos. and break if an odd no. is found
		for _, ele := range a {
			if ele%2 == 0 {
				evenNo = evenNo + 1
			} else {
				break
			}
		}
		if evenNo > 0 {
			// sort even nos. and odd nos. separately
			sort.Ints(a[0:evenNo])
			sort.Ints(a[evenNo:])
		} else {
			// either array is empty or all nos. are odd so sort
			sort.Ints(a)
		}
	}
	fmt.Println("Segregated Even Odd Nos: ", a)
}

func main() {
	a := []int{12, 34, 45, 9, 8, 90, 3}
	segregateEvenOdd(a)
}
