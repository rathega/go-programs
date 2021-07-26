/* GeeksForGeeks qn:
   Given an array Arr of N positive integers, find the missing elements (if any) in the range 0 to max of Arri.
   Input:
      N = 5
      Arr[] = {62, 8, 34, 5, 332}
   Output: 0-4 6-7 9-33 35-61 63-331
      Explanation: Elements in the range 0-4, 6-7, 9-33, 35-61 and 63-331 are not present.
*/
package main

import (
	"fmt"
)

func missingRange(a []int, low, high int) {
	fmt.Printf("Input: %v Range: %d to %d\n", a, low, high)
	// next will point to next expected element in low-high
	// init next to low
	nxt := low
	for _, ele := range a {
		// expected element found so increment next and continue
		if ele == nxt {
			nxt++
			continue
		}
		// array element is less than low boundary so go to next element
		// (happens only while start of array elements)
		if ele < nxt {
			continue
		}
		// current element is more than expected element so some elements are missing
		if ele > nxt {
			if nxt+1 == ele {
				// if only one element is missing
				fmt.Println(nxt)
			} else {
				// if a range of elements are missing
				fmt.Printf("%d - %d\n", nxt, ele-1)
			}
			// expected element next is updated to 1 more than current element
			nxt = ele + 1
		}
	}
	if nxt == high {
		// if only last element in range is missing
		fmt.Println(nxt)
	} else if nxt < high {
		// if a range of elements are missing in end
		fmt.Printf("%d - %d\n", nxt, high)
	}
}

func main() {
	a := []int{2, 4, 6}
	missingRange(a, 0, 99)
	a = []int{0, 1, 3, 50, 75}
	missingRange(a, 0, 99)
}
