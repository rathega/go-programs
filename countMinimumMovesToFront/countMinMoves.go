/* GeeksForGeeks qn:
   Given an array of size n such that array elements are in range from 1 to n. The task is to count a number of move-to-front operations to arrange items as {1, 2, 3,… n}. The move-to-front operation is to pick any item and place it in first position.
*/
package main

import "fmt"

// Start from last and see which elements need to be moved to front.
// Once an element is moved to front remaining elements automatically should move back
// -> this concept is used to find the min moves required
func minMoves(a []int) {
	fmt.Println("Input: ", a)
	// assume that entire array is in reverse order
	expEle := len(a)
	for i := len(a) - 1; i >= 0; i-- {
		// if expected element is found in last it need not be moved to front so decrement to get next expected element
		if a[i] == expEle {
			expEle--
		}
	}
	fmt.Println("Min move-to-front operations: ", expEle)
}

func main() {
	a := []int{3, 2, 1, 4}
	minMoves(a)
	a = []int{5, 7, 4, 3, 2, 6, 1}
	minMoves(a)
}
