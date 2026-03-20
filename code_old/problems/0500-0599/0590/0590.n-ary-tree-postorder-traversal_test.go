package _0590

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
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
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
