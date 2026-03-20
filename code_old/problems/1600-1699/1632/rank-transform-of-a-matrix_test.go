package _1632

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_matrixRankTransform(t *testing.T) {
	type args struct {
		matrix [][]int
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
			if got := matrixRankTransform(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixRankTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
