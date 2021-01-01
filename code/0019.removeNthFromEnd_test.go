package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *leetcode.ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *leetcode.ListNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: &leetcode.ListNode{
					Val: 1,
					Next: &leetcode.ListNode{
						Val: 2,
						Next: &leetcode.ListNode{
							Val: 3,
							Next: &leetcode.ListNode{
								Val: 4,
								Next: &leetcode.ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				n: 2,
			},
			want: &leetcode.ListNode{
				Val: 1,
				Next: &leetcode.ListNode{
					Val: 2,
					Next: &leetcode.ListNode{
						Val: 3,
						Next: &leetcode.ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "",
			args: args{
				head: &leetcode.ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !deepEqualRemoveNthFromEnd(t, got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqualRemoveNthFromEnd(t *testing.T, got, want *leetcode.ListNode) bool {
	if got == nil || want == nil {
		return got == want
	}
	if got.Val == want.Val {
		return deepEqualRemoveNthFromEnd(t, got.Next, want.Next)
	}
	t.Errorf("removeNthFromEnd() = %v, want %v", got.Val, want.Val)
	return false
}
