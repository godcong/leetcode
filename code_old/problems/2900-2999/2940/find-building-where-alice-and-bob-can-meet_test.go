package _2940

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_leftmostBuildingQueries(t *testing.T) {
	type args struct {
		heights []int
		queries [][]int
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
			if got := leftmostBuildingQueries(tt.args.heights, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leftmostBuildingQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
