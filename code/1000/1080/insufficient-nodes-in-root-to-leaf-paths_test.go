package _1080

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sufficientSubset(t *testing.T) {
	type args struct {
		root  *TreeNode
		limit int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sufficientSubset(tt.args.root, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sufficientSubset() = %v, want %v", got, tt.want)
			}
		})
	}
}
