package _1329

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_diagonalSort(t *testing.T) {
	type args struct {
		mat [][]int
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
			if got := diagonalSort(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
