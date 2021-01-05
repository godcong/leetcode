package leetcode

import (
	"testing"

	"github.com/godcong/leetcode/code"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *code.ListNode
	}
	tests := []struct {
		name string
		args args
		want *code.ListNode
	}{
		{
			name: "",
			args: args{
				node: &code.ListNode{
					Val: 1,
					Next: &code.ListNode{
						Val: 2,
						Next: &code.ListNode{
							Val: 3,
							Next: &code.ListNode{
								Val:  4,
								Next: &code.ListNode{},
							},
						},
					},
				},
			},
			want: &code.ListNode{
				Val: 3,
				Next: &code.ListNode{
					Val:  4,
					Next: &code.ListNode{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if deleteNode(tt.args.node.Next); !deepEqualDeleteNode(t, tt.args.node.Next, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", tt.args.node.Next, tt.want)
			}
		})
	}
}

func deepEqualDeleteNode(t *testing.T, got, want *code.ListNode) bool {
	if got == nil || want == nil {
		return got == want
	}
	if got.Val == want.Val {
		return deepEqualDeleteNode(t, got.Next, want.Next)
	}
	t.Errorf("deleteNode() = %v, want %v", got.Val, want.Val)
	return false
}
