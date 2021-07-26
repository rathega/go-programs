/* GeeksForGeeks qn:
   Given a sorted array A of size N, delete all the duplicates elements from A.
*/

package main

import "fmt"

func remDup(a []int) {
	fmt.Println("Input: ", a)
	// len check for empty array needed else will fail in a=a[:j+1]
	if len(a) == 0 {
		fmt.Println("Empty array")
		return
	}
	// j will hold the index of non repeating elements
	j := 0
	// i will hold the index of input with repeating elements
	for i := 0; i < len(a)-1; i++ {
		// if a non-repeating element found increment j and add in it's place
		if a[i] != a[i+1] {
			j++
			a[j] = a[i+1]
		}
	}
	// array is trimmed till j which holds the unique elements
	a = a[:j+1]
	fmt.Println("RemDup Output: ", a)
}

func main() {
	a := []int{1, 1, 2, 3, 3, 4, 4, 4}
	remDup(a)
}
