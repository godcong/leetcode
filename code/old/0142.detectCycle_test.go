package old

import (
	"testing"
)

func Test_detectCycle(t *testing.T) {
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
				head: strToListNode("[3,2,0,-4]", 1),
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !listNodeDeepEqual(t, got, tt.want, 1) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
