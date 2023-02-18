package _1237

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findSolution(t *testing.T) {
	type args struct {
		customFunction func(int, int) int
		z              int
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
			if got := findSolution(tt.args.customFunction, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
