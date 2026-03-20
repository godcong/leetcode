package _1605

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_restoreMatrix(t *testing.T) {
	type args struct {
		rowSum []int
		colSum []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreMatrix(tt.args.rowSum, tt.args.colSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
