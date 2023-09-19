package main

import "fmt"

type Sorter struct {
	array []int
}

func New(array []int) *Sorter {
	return &Sorter{array: array}
}

func (s *Sorter) sorting() []int {
	n := len(s.array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s.array[j] > s.array[j+1] {
				s.array[j], s.array[j+1] = s.array[j+1], s.array[j]
			}
		}
	}
	return s.array
}

func main() {
	sorter := New([]int{20, 2, 7, 19, 15})

	fmt.Println(sorter.sorting())
}
