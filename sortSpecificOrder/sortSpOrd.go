/* GeeksForGeeks qn:
   Given an array A of positive integers. Your task is to sort them in such a way that the first part of the array contains odd numbers sorted in descending order, rest portion contains even numbers sorted in ascending order.
*/
package main

import (
	"fmt"
	"sort"
)

// Time: O(nLogn) Space: O(n)
// Sort and make a copy of array. Use one array to copy odd, even elements in specific order to another array
func sortSpOrd1(inp []int) {
	// copy to temp array to avoid overwrite in original input
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("Input O(n) space: ", a)

	// i) sort the array and copy in another array
	// sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// copy sorted array
	temp := make([]int, len(a))
	copy(temp, a)

	// ii) traverse one array from front and place odd elements in front, even in last in another array
	// i, j are the front, last indexes of output array
	i := 0
	j := len(a) - 1
	for ti := 0; ti < len(a); ti++ {
		if temp[ti]%2 != 0 {
			// if odd place in front
			a[i] = temp[ti]
			i++
		} else {
			// if even place in last
			a[j] = temp[ti]
			j--
		}
	}
	fmt.Printf("Output: %v\n", a)
}

// Time: O(nLogn) Space: O(1)
// move odd elements to left and even elements to right then sort both parts separately
func sortSpOrd2(inp []int) {
	// copy to temp array to avoid overwrite in original input
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("Input O(1) space: ", a)

	// i) move odd elements to left, even elements to right
	i := 0
	j := len(a) - 1
	// store number of odd elements for finding odd, even part boundaries
	nOdd := 0
	for i < j {
		// find an even element in left
		for i < j && a[i]%2 != 0 {
			i++
			nOdd++
		}
		// find an odd element in right
		for i < j && a[j]%2 == 0 {
			j--
		}
		// swap if bounds are not crossed
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	// ii) sort left part, right part separately
	// sort odd part in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(a[:nOdd])))
	// sort even part in ascending order
	sort.Sort(sort.IntSlice(a[nOdd:]))
	fmt.Printf("Output: %v\n", a)
}

func main() {
	a := []int{1, 2, 3, 5, 4, 7, 10}
	sortSpOrd1(a)
	sortSpOrd2(a)
}
