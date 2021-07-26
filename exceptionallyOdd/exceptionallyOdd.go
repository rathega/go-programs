/* GeeksForGeeks qn
   Given an array of N positive integers where all numbers occur even number of times except one number which occurs odd number of times. Find the exceptional number.
*/

package main

import (
	"fmt"
)

/* Based on following XOR concepts:
   0 ^ A = A
   A ^ A = 0
   So if a number occurs odd number of times after XOR with all numbers in array only the odd element is left
*/

func exceptionallyOdd(a []int) {
	fmt.Println("Input: ", a)
	res := 0
	// XOR all elements in array
	for _, ele := range a {
		res = res ^ ele
	}
	// XOR result is the odd element
	fmt.Println("Exceptionally odd: ", res)
}
func main() {
	a := []int{1, 2, 3, 2, 3, 1, 3}
	exceptionallyOdd(a)
}
