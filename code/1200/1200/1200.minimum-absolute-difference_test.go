package _1200

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minimumAbsDifference(t *testing.T) {
	type args struct {
		arr []int
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
			if got := minimumAbsDifference(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimumAbsDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
