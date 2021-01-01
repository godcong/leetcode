package code

import (
	"reflect"
	"testing"

	"github.com/godcong/leetcode"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *leetcode.ListNode
		l2 *leetcode.ListNode
	}
	tests := []struct {
		name string
		args args
		want *leetcode.ListNode
	}{
		{
			name: "",
			args: args{
				l1: &leetcode.ListNode{
					Val: 2,
					Next: &leetcode.ListNode{
						Val: 4,
						Next: &leetcode.ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &leetcode.ListNode{
					Val: 5,
					Next: &leetcode.ListNode{
						Val: 6,
						Next: &leetcode.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &leetcode.ListNode{
				Val: 7,
				Next: &leetcode.ListNode{
					Val: 0,
					Next: &leetcode.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},

		{
			name: "",
			args: args{
				l1: &leetcode.ListNode{
					Val:  5,
					Next: nil,
				},
				l2: &leetcode.ListNode{
					Val: 5,
					Next: &leetcode.ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
			want: &leetcode.ListNode{
				Val: 0,
				Next: &leetcode.ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
