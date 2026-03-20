package _1697

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_distanceLimitedPathsExist(t *testing.T) {
	type args struct {
		n        int
		edgeList [][]int
		queries  [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceLimitedPathsExist(tt.args.n, tt.args.edgeList, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceLimitedPathsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
