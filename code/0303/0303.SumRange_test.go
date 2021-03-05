package _0303

import (
	"reflect"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums [][]int
		fns  []string
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: strToIntArrArray("[[-2, 0, 3, -5, 2, -1], [0, 2], [2, 5], [0, 5]]"),
				fns:  strToStringArray("[\"NumArray\", \"sumRange\", \"sumRange\", \"sumRange\"]"),
			},
			want: NumArray{
				sum: strToIntArray("[null, 1, -1, -3]"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var numArray NumArray
			for i, ints := range tt.args.nums[1:] {
				switch {
				case strings.Compare(tt.args.fns[i], "NumArray") > 0:
					numArray = Constructor(ints)
					t.Log(numArray)
				case strings.Compare(tt.args.fns[i], "sumRange") > 0:
					got := numArray.SumRange(ints[0], ints[1])
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Constructor() = %v, want %v", got, tt.want)
					}
				}
			}

		})
	}
}
