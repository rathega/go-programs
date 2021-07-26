package main

/* GeeksForGeeks qn:
   Given a non-empty array of integers, find the top k elements which have the highest frequency in the array. If two numbers have the same frequency then the larger number should be given preference.
*/
import (
	"fmt"
	"sort"
)

// hashmap to store frequency of elements
var m = make(map[int]int)

// byKeys implements sort.Interface for our frequency sorting
// sort.Sort of golang does sorting in O(nLogn)
type byKeys []int

func (k byKeys) Len() int {
	return len(k)
}
func (k byKeys) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}
func (k byKeys) Less(i, j int) bool {
	fi := m[k[i]]
	fj := m[k[j]]
	if fi == fj {
		// if freq of two nos. i, j are equal and i is less than j
		if i < j {
			return false
		}
		return true
	}
	// for two nos. i, j if freq of i is less than j
	if fi < fj {
		return false
	}
	// for two nos. i, j if freq of i is greater than j
	return true
}

// Time : O(n) + O(nLogn) Space: O(n)  (nLogn for sorting)
func topKElements(a []int) {
	fmt.Println("Input: ", a)

	// keys list stores all elements and later used in sorting
	keys := []int{}

	// create entries in hashmap with elements and their frequency
	for i := 0; i < len(a); i++ {
		if _, ok := m[a[i]]; ok {
			m[a[i]] = m[a[i]] + 1
		} else {
			m[a[i]] = 1
			keys = append(keys, a[i])
		}
	}
	fmt.Println("\nFreq HashMap: ", m)
	fmt.Println("\nUnsortedKeys: ", keys)
	sort.Sort(byKeys(keys))
	fmt.Println("SortedKeys: ", keys)
}

func main() {
	a := []int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5}
	topKElements(a)
}
