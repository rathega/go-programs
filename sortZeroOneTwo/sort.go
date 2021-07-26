/* GeeksForGeeks qn:
   Given an array A[] consisting 0s, 1s and 2s. The task is to write a function that sorts the given array. The functions should put all 0s first, then all 1s and all 2s in last.
*/
package main

import (
	"fmt"
)

// Time: O(n) Sorted using three passes to array
func sortNosThreePass(a []int) {
	fmt.Println("\nInput: ", a)
	l := 0
	r := len(a) - 1
	// sort all 0s to left and 1s, 2s to right
	for l < r {
		for ; l < r && a[l] == 0; l++ {
		}
		for ; l < r && (a[r] == 2 || a[r] == 1); r-- {
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
			l++
		}
	}
	// find the index where 0s end and 1s/2s start
	mid := 0
	for ; mid < len(a) && a[mid] == 0; mid++ {
	}
	// sort 1s and 2s that are present after all 0s in start
	// updated l to r has only 1s and 2s
	l = mid
	r = len(a) - 1
	for l < r {
		for ; l < r && a[l] == 1; l++ {
		}
		for ; l < r && a[r] == 2; r-- {
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
			l++
		}
	}
	fmt.Println("Output: ", a)
}

// Time: O(n) Sorted using two passes to array
func sortNosTwoPass(a []int) {
	fmt.Println("\nInput: ", a)
	zero := 0
	one := 0
	two := 0
	// count no. of 0s, 1s, 2s
	for _, ele := range a {
		if ele == 0 {
			zero++
		} else if ele == 1 {
			one++
		} else {
			two++
		}
	}
	// refill array with 0s, 1s, 2s count in order
	for i := 0; i < len(a); i++ {
		if zero > 0 {
			a[i] = 0
			zero--
		} else if one > 0 {
			a[i] = 1
			one--
		} else {
			a[i] = 2
			two--
		}
	}
	fmt.Println("Output: ", a)
}

// Time: O(n) Sorted using one pass to array
func sortNosOnePass(a []int) {
	fmt.Println("\nInput: ", a)
	// l, m, r are the indices of 0, 1, 2 respectively
	l := 0
	// start index of 1 from start and increment as we find 1s in array
	m := 0
	r := len(a) - 1
	for m <= r {
		if a[m] == 0 {
			// 0 is found so increment 0s index
			// m is also incremented since it's index starts from beginning
			a[l], a[m] = a[m], a[l]
			l++
			m++
		} else if a[m] == 1 {
			// 1 is found so increment 1s index
			m++
		} else {
			// 2 is found so decrement 2s index
			a[r], a[m] = a[m], a[r]
			r--
		}
	}
	fmt.Println("Output: ", a)
}

func main() {
	a := []int{0, 2, 1, 2, 0}
	sortNosThreePass(a)
	sortNosTwoPass(a)
	sortNosOnePass(a)
	a = []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	sortNosThreePass(a)
	sortNosTwoPass(a)
	sortNosOnePass(a)
}
