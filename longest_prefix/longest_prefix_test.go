package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_longestPrefix(t *testing.T) {

	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Correct test",
			args: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Correct test",
			args: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Correct test",
			args: []string{"rocket", "racecar", "rise"},
			want: "r",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.args)
			require.Equal(t, got, tt.want)
		})
	}
}
