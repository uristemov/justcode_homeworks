package main

import "testing"
import "github.com/stretchr/testify/require"

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Correct test",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "Correct test",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "Uncorrect test",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 20,
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.args.nums, tt.args.target)
			require.Equal(t, got, tt.want)
		})
	}
}
