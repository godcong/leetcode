package leetcode

import (
	"reflect"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
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
				nums: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallerNumbersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
