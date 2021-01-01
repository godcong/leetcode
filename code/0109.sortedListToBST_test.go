package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *leetcode.ListNode
	}
	tests := []struct {
		name string
		args args
		want *leetcode.TreeNode
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
