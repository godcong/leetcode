package _0459

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_find132pattern(t *testing.T) {
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
				nums: StrToIntArray("[1,2,3,4]"),
			},
			want: false,
		},
		{
			name: "",
			args: args{
				nums: StrToIntArray("[3,1,4,2]"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
