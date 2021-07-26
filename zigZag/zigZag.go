/* GeeksForGeeks qn:
Given an array Arr (distinct elements) of size N. Rearrange the elements of array in zig-zag fashion. The converted array should be in form a < b > c < d > e < f. The relative order of elements is same in the output i.e you have to iterate on the original array only.
Input:
N = 7
Arr[] = {4, 3, 7, 8, 6, 2, 1}
Output: 3 7 4 8 2 6 1
Explanation: 3 < 7 > 4 < 8 > 2 < 6 > 1
*/

package main

import (
	"fmt"
)

func zigZagArray(a []int) {
	fmt.Println("Input: ", a)
	for i := 0; i < len(a)-1; i++ {
		if i%2 == 0 {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		} else {
			if a[i] < a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
	fmt.Println("ZigZag: ", a)
}

func main() {
	a := []int{4, 3, 7, 8, 6, 2, 1}
	zigZagArray(a)
}
