package _0927

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_threeEqualParts(t *testing.T) {
	type args struct {
		arr []int
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
			if got := threeEqualParts(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeEqualParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
