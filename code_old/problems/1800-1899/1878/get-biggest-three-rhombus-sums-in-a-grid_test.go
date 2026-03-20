package _1878

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getBiggestThree(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := getBiggestThree(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBiggestThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
