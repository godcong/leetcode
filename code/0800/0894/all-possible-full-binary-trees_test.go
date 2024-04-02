package _0894

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
