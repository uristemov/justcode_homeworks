package main

import "testing"
import "github.com/stretchr/testify/require"

func Test_compareTwoSlices(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Two equal slices",
			args: args{
				a: []int{2, 4, 5, 8, 10, 15, 5},
				b: []int{2, 4, 5, 8, 10, 15, 5},
			},
			want: true,
		},
		{
			name: "Two not equal slices",
			args: args{
				a: []int{2, 4, 5, 8, 10, 15, 5},
				b: []int{2, 4, 5, 8, 10, 15, 5, 6, 7},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareTwoSlices(tt.args.a, tt.args.b)
			require.Equal(t, got, tt.want)
		})
	}
}

func Test_compareTwoSlicesWithOrder(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Two not equal slices",
			args: args{
				a: []int{2, 4, 5, 8, 10, 15, 5, 5, 7},
				b: []int{2, 4, 5, 8, 10, 15, 5, 5, 7, 7},
			},
			want: false,
		},
		{
			name: "Two not equal slices",
			args: args{
				a: []int{2, 4, 5, 8, 10, 15, 5, 5, 6, 7},
				b: []int{2, 4, 5, 8, 10, 15, 5, 6, 7},
			},
			want: false,
		},
		{
			name: "Two equal slices",
			args: args{
				a: []int{2, 4, 5, 8, 10, 15, 5, 5, 6, 7},
				b: []int{5, 2, 4, 5, 8, 10, 15, 5, 6, 7},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareTwoSlicesWithOrder(tt.args.a, tt.args.b)
			require.Equal(t, got, tt.want)
		})
	}
}
