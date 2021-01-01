package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_isPossible(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: leetcode.strToIntArray("[1,2,3,3,4,5]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: leetcode.strToIntArray("[1,2,3,3,4,4,5,5]"),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: leetcode.strToIntArray("[1,2,3,4,4,5]"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossible(tt.args.nums); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
