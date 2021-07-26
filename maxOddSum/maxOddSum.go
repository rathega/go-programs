/* GeeksForGeeks qn:
   Given an array of integers, check whether there is a subsequence with odd sum and if yes, then finding the maximum odd sum. If no subsequence contains odd sum, print -1.
*/
package main

import (
	"fmt"
	"math"
)

func maxOdd(a []int) {
	fmt.Println("Input: ", a)

	// sum holds the max sum
	sum := 0
	// negOdd holds the largest negative odd
	negOdd := math.MinInt32
	// posOdd holds the smallest positive odd
	posOdd := math.MaxInt32
	// flag to indicate if atleast one odd number is present
	oddFound := false
	for _, ele := range a {
		// add all positive numbers to sum
		if ele > 0 {
			sum = sum + ele
		}
		// update flags if an odd number comes
		if ele%2 != 0 {
			oddFound = true
			if ele < 0 && negOdd < ele {
				negOdd = ele
			}
			if ele > 0 && posOdd > ele {
				posOdd = ele
			}
		}
	}
	// if sum is already negative that is the largest
	if sum%2 != 0 {
		fmt.Println(sum)
		return
	}
	// if not even one odd is found then odd sum not possible
	if !oddFound {
		fmt.Println("-1")
		return
	}
	// check if removing positive odd or adding negative odd gives maximum and return
	if sum-posOdd > sum+negOdd {
		fmt.Println(sum - posOdd)
	} else {
		fmt.Println(sum + negOdd)
	}
}

func main() {
	maxOdd([]int{2, 5, -4, 3, -1})
	maxOdd([]int{4, -3, 3, -5})
	maxOdd([]int{2, 4, 6})
}
