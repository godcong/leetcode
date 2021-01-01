package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_minPatches(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: leetcode.strToIntArray("[1,3]"),
				n:    6,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: leetcode.strToIntArray("[1,5,10]"),
				n:    20,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatches(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
