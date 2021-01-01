package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		t1 *leetcode.TreeNode
		t2 *leetcode.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *leetcode.TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.t1, tt.args.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
