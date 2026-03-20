package _1073

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_addNegabinary(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
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
			if got := addNegabinary(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addNegabinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
