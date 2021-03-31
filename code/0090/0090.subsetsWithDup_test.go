package _0090

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: StrToIntArray("[1,2,2]"),
			},
			want: StrToIntArrArray("[[],[1],[1,2],[1,2,2],[2],[2,2]]"),
		},
		{
			name: "",
			args: args{
				nums: StrToIntArray("[0]"),
			},
			want: StrToIntArrArray("[[],[0]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetsWithDup(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
