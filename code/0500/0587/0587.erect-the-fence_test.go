package _0587

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_outerTrees(t *testing.T) {
	type args struct {
		trees [][]int
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
			if got := outerTrees(tt.args.trees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
