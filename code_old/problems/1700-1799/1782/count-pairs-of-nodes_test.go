package _1782

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		n       int
		edges   [][]int
		queries []int
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
			if got := countPairs(tt.args.n, tt.args.edges, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
