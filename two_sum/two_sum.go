package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for idx, value := range nums {
		if index, exists := m[target-value]; exists {
			return []int{index, idx}
		}

		m[value] = idx
	}
	return []int{}
}
