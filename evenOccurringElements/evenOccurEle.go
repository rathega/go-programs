/* GeeksForGeeks qn:
   Given an array Arr of N integers that contains odd number of occurrences for all numbers except for a few elements which are present even number of times. Find the elements which have even occurrences in the array.
*/
package main

import (
	"fmt"
	"strconv"
)

// Time: O(n) Space: O(1)
// for each element in array left shift 1 to that element and xor to result
//   -> so each bit of result is made to represent each element (1<<ele)
//   -> if a no. occurs even times that bit of result will be 0 (X ^ X = 0)
//   -> if a no. occurs odd times that bit of result will be 1 (0 ^ X = X)
func evenOccurNos(a []int) {
	res := 0
	for _, ele := range a {
		res = res ^ 1<<ele
		fmt.Printf("after adding %d res: %s\n", ele, strconv.FormatInt(int64(res), 2))
	}
	fmt.Println("\nEven occurring nos. are: ")
	for _, ele := range a {
		if 1<<ele&res != 1<<ele { // for odd occurring elements change to 1<<ele&res == 1<<ele
			fmt.Println(ele)
			res = res ^ 1<<ele
		}
	}
}
func main() {
	a := []int{9, 12, 23, 10, 12, 12, 15, 23, 14, 12, 15}
	evenOccurNos(a)
}
