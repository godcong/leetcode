package _2961

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getGoodIndices(t *testing.T) {
	type args struct {
		variables [][]int
		target    int
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
			if got := getGoodIndices(tt.args.variables, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGoodIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
