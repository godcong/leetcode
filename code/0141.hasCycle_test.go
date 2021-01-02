package code

import (
	"testing"

	"github.com/godcong/leetcode"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: leetcode.strToListNode("[3,2,0,-4]", 1),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: leetcode.strToListNode("[1,2]", 0),
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: leetcode.strToListNode("[1]", -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
