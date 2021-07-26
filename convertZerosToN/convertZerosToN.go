/* GeeksForGeeks qn:
   You are given an integer N. You need to convert all zeroes of N to 5.
*/

package main

import (
	"fmt"
)

// convert all 0s in number a to x
func convertZerosToN(a int, x int) {
	temp := a
	fmt.Println(temp)
	// loop till each digit is accessed
	for i := 1; a > 0; i, a = i*10, a/10 {
		// read each digit
		digit := a % 10
		// if digit is 0 then add x/10x/i*x to number based on digit in ones/tens/i's place and so on
		if digit == 0 {
			temp = temp + (i * x)
		}
	}
	fmt.Println(temp)
}

func main() {
	convertZerosToN(1100, 9)
}
