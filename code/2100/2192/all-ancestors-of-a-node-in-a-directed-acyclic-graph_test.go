package _2192

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getAncestors(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
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
			if got := getAncestors(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAncestors() = %v, want %v", got, tt.want)
			}
		})
	}
}
