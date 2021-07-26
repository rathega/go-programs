/* GeeksForGeeks qn:
   Given an array of numbers form 0 to 9 of size N. Your task is to rearrange elements of the array such that after combining all the elements of the array number formed is maximum.
*/
package main

import (
	"fmt"
	"strconv"
)

func largestNo(a []int) {
	fmt.Println("Input: ", a)

	// temp to store count of occurrences of digits 0-9
	temp := make([]int, 10)
	// count the no. of occurrences of each digit
	for _, ele := range a {
		temp[ele] = temp[ele] + 1
	}
	// initialise result to empty string
	res := ""
	// run loop from 9 to 0 and append to result
	for i := 9; i >= 0; i-- {
		// append a no. x times to the result for x times in input
		for temp[i] != 0 {
			res = res + strconv.Itoa(i)
			temp[i] = temp[i] - 1
		}
	}

	fmt.Println("LargesNo: ", res)
}

func main() {
	a := []int{9, 0, 1, 3, 0}
	largestNo(a)
}
