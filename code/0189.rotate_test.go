package code

import (
	"reflect"
	"testing"
)

func Test_rotate0189(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: strToIntArray("[1,2,3,4,5,6,7]"),
				k:    3,
			},
			want: strToIntArray("[5,6,7,1,2,3,4]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate0189(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("maxProfit0189() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
