/* GeeksForGeeks qn:
   Given an array of integers. Check whether it contains a triplet that sums up to zero.
*/
package main

import (
	"fmt"
	"sort" // required for tripletSum3 approach
)

// O(n*n*n)
// iterate and compare each combination of three no. to find the sum
// Extras:-
//   outer loop 'for' condition 'i<len(a)-2' -> i<len(a) gives same result but avoiding since middle loop won't proceed that iteration
//   middle loop 'for' condition 'i<len(a)-1' -> i<len(a) gives same result but avoiding since inner loop won't proceed that iteration
func tripletSum1(a []int, sum int) {
	fmt.Println("Approach n*n*n: ", a)
	for i := 0; i < len(a)-2; i++ {
		for j := i + 1; j < len(a)-1; j++ {
			for k := j + 1; k < len(a); k++ {
				if a[i]+a[j]+a[k] == 0 {
					fmt.Printf("%d %d %d\n", a[i], a[j], a[k])
				}
			}
		}
	}
}

// O(n*n) with extra space
// reduce one loop by storing elements in hash map for each outer iteration
// and clear hash map for each outer loop iteration to avoid duplicate results
// Extras:-
//   outer loop 'for' condition 'i<len(a)-1' -> i<len(a) gives same result but avoiding since inner loop won't proceed that iteration
func tripletSum2(a []int, sum int) {
	fmt.Println("Approach n*n extra space: ", a)
	h := make(map[int]struct{})
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			// considering ith, jth elements what is the third element to arrive at required sum
			rem := sum - (a[i] + a[j])
			// check if third element exist if not add it to hash map
			if _, ok := h[rem]; ok {
				fmt.Printf("%d %d %d\n", a[i], a[j], rem)
			} else {
				h[a[j]] = struct{}{}
			}
		}
		// clear hash map for every outer loop iteration -> otherwise will give duplicate results eg: 2,-3,1 and -3,1,2
		h = make(map[int]struct{})
	}
}

// O(n*Logn) + O(n*n) = O(n*n) without extra space
// sort the array and apply 'pairSum' approach for each of the element
// Extras:-
//   outer loop 'for' condition 'i<len(a)-2' -> i<len(a) gives same result but avoiding since inner pairSum won't proceed that iteration
//   pairSum - in a sorted consider first, last element as pair that equals sum and move left or right based on sum is >/< to sum
func tripletSum3(a []int, sum int) {
	fmt.Println("Approach n*n no extra space: ", a)
	sort.Ints(a)
	for k := 0; k < len(a)-2; k++ {
		// pairSum algo for remaining elements
		// remaining value required to form triplet with current element taken as 'sum' for pairSum algo
		i := k + 1
		j := len(a) - 1
		rem := sum - a[k]
		for i < j {
			if a[i]+a[j] == rem {
				fmt.Printf("%d %d %d\n", a[k], a[i], a[j])
				i++
				j--
			} else if a[i]+a[j] > rem {
				j--
			} else {
				i++
			}
		}
	}
}

func main() {
	a := []int{0, -1, 2, -3, 1}
	tripletSum1(a, 0)
	tripletSum2(a, 0)
	tripletSum3(a, 0)
}
