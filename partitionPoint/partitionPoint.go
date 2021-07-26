/* GeeksForGeeks qn:
   Given an unsorted array of integers. Find an element such that all the elements to its left are smaller and to its right are greater. Print -1 if no such element exists
   Note that there can be more than one such element. In that case print the first such number occurring in the array.
*/
package main

import (
	"fmt"
)

// Time: O(n) Space: O(1)
// Idea is i) if there is no partition element and an element > current maximum is found assume greater element as partition element
//         ii) if partition element is set and found an element <= partition element then unset the current partition element
//         In both cases update currentMax with current maximum element in array
func partitionPoint(inp []int) {
	// copy input to avoid overwrite on original input
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("Input: ", a)

	// init partitionElement and currentMax to -1
	partEle := -1
	curMax := -1
	for i := 0; i < len(a); i++ {
		if partEle == -1 && a[i] > curMax {
			// case i)
			partEle = a[i]
		} else if partEle != -1 && a[i] <= partEle {
			// case ii)
			partEle = -1
		}
		// update currentMax
		if curMax < a[i] {
			curMax = a[i]
		}
	}
	// at the end of for loop partition element if found will be present in partEle or -1 if none present
	fmt.Println("Parition element: ", partEle)
}

func main() {
	fmt.Println("Hello, playground")
	a := []int{4, 3, 2, 5, 8, 6, 7}
	partitionPoint(a)
	a = []int{2823, 2301, 486, 2737, 851, 689, 688, 515, 196, 1522, 1571, 1645, 2729, 2153, 2021, 2556, 1493, 1591, 2229, 692, 1102, 2770, 2719, 2720, 2983, 2869, 820, 259, 2862}
	partitionPoint(a)
}
