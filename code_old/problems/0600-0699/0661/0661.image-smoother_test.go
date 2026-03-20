package _0661

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_imageSmoother(t *testing.T) {
	type args struct {
		img [][]int
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
			if got := imageSmoother(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageSmoother() = %v, want %v", got, tt.want)
			}
		})
	}
}
