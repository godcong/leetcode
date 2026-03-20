package _3379

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_constructTransformedArray(t *testing.T) {
	type args struct {
		nums []int
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
			if got := constructTransformedArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructTransformedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
