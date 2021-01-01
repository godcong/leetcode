package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:   leetcode.strToIntArray("[5,7,7,8,8,10]"),
				target: 8,
			},
			want: leetcode.strToIntArray("[3,4]"),
		},
		{
			name: "",
			args: args{
				nums:   leetcode.strToIntArray("[5,7,7,8,8,10]"),
				target: 6,
			},
			want: leetcode.strToIntArray("[-1,-1]"),
		},
		{
			name: "",
			args: args{
				nums:   leetcode.strToIntArray("[]"),
				target: 0,
			},
			want: leetcode.strToIntArray("[-1,-1]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
