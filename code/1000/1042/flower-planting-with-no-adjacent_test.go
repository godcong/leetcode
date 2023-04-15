package _1042

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_gardenNoAdj(t *testing.T) {
	type args struct {
		n     int
		paths [][]int
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
			if got := gardenNoAdj(tt.args.n, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gardenNoAdj() = %v, want %v", got, tt.want)
			}
		})
	}
}
