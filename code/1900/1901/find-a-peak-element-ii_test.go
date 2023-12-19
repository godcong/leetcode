package _1901

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findPeakGrid(t *testing.T) {
	type args struct {
		mat [][]int
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
			if got := findPeakGrid(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeakGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
