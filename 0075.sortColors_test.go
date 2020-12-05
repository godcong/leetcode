package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{
					2, 0, 2, 1, 1, 0,
				},
			},
			want: []int{
				0, 0, 1, 1, 2, 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("climbStairs() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
