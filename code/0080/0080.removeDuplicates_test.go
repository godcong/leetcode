package _0080

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
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
				nums: StrToIntArray("[1,1,1,2,2,3]"),
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				nums: StrToIntArray("[0,0,1,1,1,1,2,3,3]"),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
