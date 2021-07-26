/* GeeksForGeeks qn:
   Given a sorted array of N integers, find the number of distinct absolute values among the elements of the array.
   Input:
        N = 6
        arr[] = {-3, -2, 0, 3, 4, 5}
   Output: 5
        Explanation: There are 5 distinct absolute
        values among the elements of this array,
        i.e. 0, 2, 3, 4 and 5.
*/
package main

import "fmt"

// Space: O(n) Time: O(n)
// Create a map
// For positive elements add entry if not in map
// For negative elements negate and add entry if not in map
// No. of elements in map gives the absolute elements count i.e. size
func absDistCount1(inp []int) {
	// copy input to avoid overwrite on original input
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("\nInput: ", a)

	// create map to add distinct elements
	m := make(map[int]struct{})
	for _, ele := range a {
		// if element is negative make it positive
		if ele < 0 {
			ele = -1 * ele
		}
		// add entry if not present in map
		if _, ok := m[ele]; !ok {
			m[ele] = struct{}{}
		}
	}
	fmt.Println("Space O(n) Output: ", len(m))
}

// Space: O(1) Time: O(n)
// Init two pointers one on front and one on last; Init count to length of array
// Traverse front from left to right if adjacent elements are same count--
// Traverse last from right to left if adjacent elements are same count--
// If addition of two elements are same count--
func absDistCount2(inp []int) {
	// copy input to avoid overwrite on original input
	a := make([]int, len(inp))
	copy(a, inp)

	fmt.Println("Input: ", a)

	// assume all elements are distinct and init count to length of array
	count := len(a)
	for i, j := 0, len(a)-1; i < j; {
		// checks for same elements on front
		for ; i < j && a[i] == a[i+1]; i++ {
			count--
		}
		// checks for same elements on last
		for ; i < j && a[j] == a[j-1]; j-- {
			count--
		}
		// if front and last cross all elements are traversed so break
		if i >= j {
			break
		}
		if a[i]+a[j] == 0 {
			// if sum is 0 one element is negative of other so decrement count
			i++
			j--
			count--
		} else if a[i]+a[j] > 0 {
			// if sum is >0 move last index
			j--
		} else {
			// if sum is <0 move front index
			i++
		}
	}
	fmt.Println("Space O(1)) Output: ", count)
}

func main() {
	a := []int{-3, -2, 0, 3, 4, 5}
	absDistCount1(a)
	absDistCount2(a)
	a = []int{-1, -1, -1, -1, 0, 1, 1, 1, 1}
	absDistCount1(a)
	absDistCount2(a)
	a = []int{0, 0, 0}
	absDistCount1(a)
	absDistCount2(a)
}
