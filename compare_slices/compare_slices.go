// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	a := []int{2, 4, 5, 8, 10, 15, 5}
	b := []int{2, 4, 5, 8, 10, 15, 5}

	fmt.Println(compareTwoSlices(a, b))
	fmt.Println(compareTwoSlicesWithOrder(a, b))

}

func compareTwoSlicesWithOrder(a, b []int) bool {
	aMap := make(map[int]int, len(a))
	bMap := make(map[int]int, len(b))

	for _, aVal := range a {
		aMap[aVal]++
	}

	for _, bVal := range b {
		bMap[bVal]++
	}
	//fmt.Println(aMap)
	//fmt.Println(bMap)
	for aKey, aVal := range aMap {
		if bMap[aKey] != aVal {
			return false
		}
	}
	return true

}

func compareTwoSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, value := range a {
		if value != b[idx] {
			return false
		}
	}
	return true
}
