package _0092

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
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
				head:  StrToListNode("[1,2,3,4,5]", -1),
				left:  2,
				right: 4,
			},
			want: StrToListNode("[1,4,3,2,5]", -1),
		},
		{
			name: "",
			args: args{
				head:  StrToListNode("[5]", -1),
				left:  1,
				right: 1,
			},
			want: StrToListNode("[5]", -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
