package _2003

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_smallestMissingValueSubtree(t *testing.T) {
	type args struct {
		parents []int
		nums    []int
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
			if got := smallestMissingValueSubtree(tt.args.parents, tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestMissingValueSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
