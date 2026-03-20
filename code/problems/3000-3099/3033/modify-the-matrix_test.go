package _3033

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_modifiedMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
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
			if got := modifiedMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
