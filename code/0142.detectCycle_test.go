package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *leetcode.ListNode
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
				head: leetcode.strToListNode("[3,2,0,-4]", 1),
			},
			want: &leetcode.ListNode{
				Val: 2,
				Next: &leetcode.ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !leetcode.listNodeDeepEqual(t, got, tt.want, 1) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
