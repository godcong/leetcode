package _2965

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findMissingAndRepeatedValues(t *testing.T) {
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
			if got := findMissingAndRepeatedValues(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMissingAndRepeatedValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
