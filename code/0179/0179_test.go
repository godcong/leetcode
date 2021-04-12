package _0179

import (
	. "github.com/godcong/leetcode/common"
	"testing"
)

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: StrToIntArray(" [10,2]"),
			},
			want: "210",
		},
		{
			name: "",
			args: args{
				nums: StrToIntArray("[3,30,34,5,9]"),
			},
			want: "9534330",
		},
		{
			name: "",
			args: args{
				nums: StrToIntArray("[1]"),
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
