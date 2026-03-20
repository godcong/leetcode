package _3112

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		n         int
		edges     [][]int
		disappear []int
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
			if got := minimumTime(tt.args.n, tt.args.edges, tt.args.disappear); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
