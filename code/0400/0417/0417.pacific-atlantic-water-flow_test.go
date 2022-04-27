package _0417

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
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
			if got := pacificAtlantic(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pacificAtlantic() = %v, want %v", got, tt.want)
			}
		})
	}
}
