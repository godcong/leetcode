package _LCP_24

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numsGame(t *testing.T) {
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
			if got := numsGame(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numsGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
