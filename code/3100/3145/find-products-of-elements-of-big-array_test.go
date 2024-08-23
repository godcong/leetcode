package _3145

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findProductsOfElements(t *testing.T) {
	type args struct {
		queries [][]int64
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
			if got := findProductsOfElements(tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findProductsOfElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
