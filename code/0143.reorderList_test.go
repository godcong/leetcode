package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: strToListNode("[1,2,3,4]", -1),
			},
			want: strToListNode("[1,4,2,3]", -1),
		},
		{
			name: "",
			args: args{
				head: strToListNode("1->2->3->4->5", -1),
			},
			want: strToListNode("1->5->2->4->3", -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
			if !deepEqualReorderList(t, tt.args.head, tt.want) {
				t.Errorf("reorderList() = %v, want %v", tt.args.head, tt.want)
			}
		})
	}
}

func deepEqualReorderList(t *testing.T, got, want *ListNode) bool {
	if got == nil || want == nil {
		return got == want
	}
	if got.Val == want.Val {
		return deepEqualDeleteNode(t, got.Next, want.Next)
	}
	t.Errorf("reorderList() = %v, want %v", got.Val, want.Val)
	return false
}
