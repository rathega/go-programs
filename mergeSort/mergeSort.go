package main

import (
	"fmt"
)

// merge function used in mergeSort
func merge(a []int, l, mid, r int) {
	// temporary array to store sorted array
	temp := make([]int, r-l+1)
	// i is starting of first half, j is starting of second half, k is starting of temp array
	var i, j, k int
	for i, j, k = l, mid+1, 0; i <= mid && j <= r; k++ {
		if a[i] < a[j] {
			temp[k] = a[i]
			i++
		} else {
			temp[k] = a[j]
			j++
		}
	}
	if j > r {
		// copy remaining elements of first half
		for ; i <= mid; i, k = i+1, k+1 {
			temp[k] = a[i]
		}
	} else {
		// copy remaining elements of second half
		for ; j <= r; j, k = j+1, k+1 {
			temp[k] = a[j]
		}
	}
	// store sorted elements back to original array
	for i, j = 0, l; i < len(temp); i, j = i+1, j+1 {
		a[j] = temp[i]
	}
}

// Time: O(nLogn) Space: O(n) space is used to store temporary sorted results
// mergeSort function
func mergeSort(a []int, l, r int) {
	// ending condition for recursion
	if l >= r {
		return
	}
	// find mid to split array of bounds l and r
	mid := l + (r-l)/2
	// sort first half and second half separately
	mergeSort(a, l, mid)
	mergeSort(a, mid+1, r)
	// merge sorted halves to get fully sorted array
	merge(a, l, mid, r)
}

func main() {
	a := []int{5, 4, 3, 2, 1}
	fmt.Println("Input: ", a)
	mergeSort(a, 0, len(a)-1)
	fmt.Println("Output: ", a)
}
