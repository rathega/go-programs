/* GeeksForGeeks qn:
   Given a array, write a program to construct a triangle where last row contains elements of given array, every element of second last row contains sum of below two elements and so on.
   Input:
      A[] = {4, 7, 3, 6, 7};
   Output:
      81 40 41 21 19 22 11 10 9 13 4 7 3 6 7
   Explanation:
          81
        40  41
      21  19  22
    11  10   9   13
   4   7   3   6    7
*/
package main

import "fmt"

// Without recursion
func printTriangle(inp []int) {
	// copy input to avoid overwrite on original input
	a := make([]int, len(inp))
	copy(a, inp)

	// a: [4 7 3 6 7]

	// store original size of input
	alen := len(a)
	// res holds the original element plus next tier summed elements
	res := []int{}
	// first append original elements to res
	res = append(res, a...)
	// res: [4 7 3 6 7]

	// for loop to compute sum of each tier and append to res
	for n := alen; n > 0; n-- {
		for i := 0; i < n-1; i = i + 1 {
			res = append(res, a[i]+a[i+1])
			a[i] = a[i] + a[i+1]
		}
	}
	// res: [4 7 3 6 7 11 10 9 13 21 19 22 40 41 81]

	// iterate throught res from end and print n elements from end in each iteration
	for n := 1; n <= alen; n++ {
		i := len(res) - n
		fmt.Println(res[i:])
		// printed n elements are removed from res
		res = res[:len(res)-n]
	}
}

// With recursion
func recPrintTriangle(a []int, n int) {
	if n == 0 {
		return
	}
	b := []int{}
	for i := 0; i < n-1; i++ {
		b = append(b, a[i]+a[i+1])
	}
	recPrintTriangle(b, len(b))
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}

func main() {
	a := []int{4, 7, 3, 6, 7}
	fmt.Println("Input: ", a)
	fmt.Println("\nWithout recursion: ")
	printTriangle(a)
	fmt.Println("\nWith recursion: ")
	recPrintTriangle(a, len(a))
}
