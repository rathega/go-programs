/*
GeeksForGeeks qn:
Given a number N. Your task is to check whether it is fascinating or not.
Fascinating Number: When a number(should contain 3 digits or more) is multiplied by 2 and 3 ,and when both these products are concatenated with the original number, then it results in all digits from 1 to 9 present exactly once.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isFascinating(n int) bool {
	fmt.Println("\nInput: ", n)
	s := strconv.Itoa(n)
	if len(s) < 3 {
		return false
	}
	s = strconv.Itoa(n*2) + strconv.Itoa(n*3) + strconv.Itoa(n)
	fmt.Println("Concatenated value: ", s)
	if len(s) != 9 {
		return false
	}
	for i := 1; i <= 9; i++ {
		s = strings.Replace(s, strconv.Itoa(i), "", 1)
	}
	if len(s) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isFascinating(192))
	fmt.Println(isFascinating(123))
}
