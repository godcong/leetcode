package leetcode

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums  []int
		after []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums:  strToIntArray("1,2,3"),
				after: strToIntArray("1,3,2"),
			},
		},
		{
			name: "",
			args: args{
				nums:  strToIntArray("3,2,1"),
				after: strToIntArray("1,2,3"),
			},
		},
		{
			name: "",
			args: args{
				nums:  strToIntArray("1,1,5"),
				after: strToIntArray("1,5,1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.args.after) {
				t.Errorf("want %v get %v", tt.args.after, tt.args.nums)
			}
		})
	}
}
