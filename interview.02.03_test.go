package leetcode

import "testing"

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				node: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{},
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

func deepEqualDeleteNode(t *testing.T, got, want *ListNode) bool {
	if got == nil || want == nil {
		return got == want
	}
	if got.Val == want.Val {
		return deepEqualDeleteNode(t, got.Next, want.Next)
	}
	t.Errorf("deleteNode() = %v, want %v", got.Val, want.Val)
	return false
}
