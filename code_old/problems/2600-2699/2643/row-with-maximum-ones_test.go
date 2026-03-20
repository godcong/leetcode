package _2643

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rowAndMaximumOnes(t *testing.T) {
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
			if got := rowAndMaximumOnes(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rowAndMaximumOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
