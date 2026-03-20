package _3372

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maxTargetNodes(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
		k      int
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
			if got := maxTargetNodes(tt.args.edges1, tt.args.edges2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxTargetNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
