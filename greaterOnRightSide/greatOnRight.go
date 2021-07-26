/* GeeksForGeeks qn:
   You are given an array Arr of size N. Replace every element with the next greatest element (greatest element on its right side) in the array. Also, since there is no element next to the last element, replace it with -1.
   Input:
      N = 6
      Arr[] = {16, 17, 4, 3, 5, 2}
   Output:
      17 5 5 5 2 -1
*/
package main

import "fmt"

func greatOnRight(a []int) {
	fmt.Println("Input: ", a)
	// initialise current greatest with -1 to fill as last element
	curGreat := -1
	for i := len(a) - 1; i >= 0; i-- {
		// store temp to update current greatest and update current element
		temp := a[i]
		a[i] = curGreat
		// check and update current greatest for next iteration
		if temp > curGreat {
			curGreat = temp
		}
	}
	fmt.Println("Greater on right: ", a)
}

func main() {
	a := []int{16, 17, 4, 3, 5, 2}
	greatOnRight(a)
}
