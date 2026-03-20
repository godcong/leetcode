package _1766

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getCoprimes(t *testing.T) {
	type args struct {
		nums  []int
		edges [][]int
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
			if got := getCoprimes(tt.args.nums, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
