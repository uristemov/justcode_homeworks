package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/require"

func Test_sliceSorting(t *testing.T) {
	tests := []struct {
		name   string
		sorter *Sorter
		want   []int
	}{
		{
			name:   "Correct test",
			sorter: New([]int{3, 2, 4}),
			want:   []int{2, 3, 4},
		},
		{
			name:   "Correct test",
			sorter: New([]int{19, 1, 9, 2, 1, 18, 15}),
			want:   []int{1, 1, 2, 9, 15, 18, 19},
		},
		{
			name:   "Correct test",
			sorter: New([]int{20, 2, 7, 19, 15}),
			want:   []int{2, 7, 15, 19, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sorter.sorting()
			fmt.Println(got, "\n", tt.sorter.sorting())
			require.Equal(t, tt.want, got)
		})
	}
}
