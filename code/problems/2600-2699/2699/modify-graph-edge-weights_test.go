package _2699

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_modifiedGraphEdges(t *testing.T) {
	type args struct {
		n           int
		edges       [][]int
		source      int
		destination int
		target      int
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
			if got := modifiedGraphEdges(tt.args.n, tt.args.edges, tt.args.source, tt.args.destination, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedGraphEdges() = %v, want %v", got, tt.want)
			}
		})
	}
}
