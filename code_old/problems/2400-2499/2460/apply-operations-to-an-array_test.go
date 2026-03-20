package _2460

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_applyOperations(t *testing.T) {
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
			if got := applyOperations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
