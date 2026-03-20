package _2846

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minOperationsQueries(t *testing.T) {
	type args struct {
		n       int
		edges   [][]int
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
			if got := minOperationsQueries(tt.args.n, tt.args.edges, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperationsQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
