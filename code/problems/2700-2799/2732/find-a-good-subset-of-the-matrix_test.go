package _2732

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_goodSubsetofBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := goodSubsetofBinaryMatrix(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodSubsetofBinaryMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
