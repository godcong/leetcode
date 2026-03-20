package _2670

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distinctDifferenceArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distinctDifferenceArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distinctDifferenceArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
