package _2639

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findColumnWidth(t *testing.T) {
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
			if got := findColumnWidth(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findColumnWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}
