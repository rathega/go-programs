/* GeeksForGeeks qn:
   Given an array of integers, and a number ‘sum’, print all pairs in the array whose sum is equal to ‘sum’.
*/
package main

import (
	"fmt"
	"sort"
)

// O(n*n)
// iterate and compare each combination of two no. to find the sum
func pairSum1(a []int, sum int) {
	fmt.Println("Approach n*n: ", a)
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == sum {
				fmt.Printf("%d %d\n", a[i], a[j])
			}
		}
	}
}

// O(n) with extra space
// reduce one loop by storing elements in hash map
// Note: *this method can be extended to duplicate elements by storing no. of occurrence of element in hash map[int]int
//       *in if ok check print same a[i], j equal to no. of times stored in hash map
func pairSum2(a []int, sum int) {
	fmt.Println("Approach n extra space: ", a)
	h := make(map[int]struct{})
	for i := 0; i < len(a); i++ {
		j := sum - a[i]
		if _, ok := h[j]; ok {
			fmt.Printf("%d %d\n", a[i], j)
		} else {
			h[a[i]] = struct{}{}
		}
	}
}

// O(n*Logn) + O(n) = O(n*Logn) without extra space
// sort array then consider first, last elements as pair that equals sum and move left or right based on sum is >/< to sum
func pairSum3(a []int, sum int) {
	fmt.Println("Approach n no extra space: ", a)
	sort.Ints(a)
	i := 0
	j := len(a) - 1
	for i < j {
		if a[i]+a[j] == sum {
			fmt.Printf("%d %d\n", a[i], a[j])
			i++
			j--
		} else if a[i]+a[j] > sum {
			j--
		} else {
			i++
		}
	}
}

func main() {
	a := []int{1, 5, 7, -1, 6}
	pairSum1(a, 6)
	pairSum2(a, 6)
	pairSum3(a, 6)
}
