package _0061

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				head: StrToListNode("[1,2,3,4,5]", -1),
				k:    2,
			},
			want: StrToListNode("[4,5,1,2,3]", -1),
		},
		{
			name: "",
			args: args{
				head: StrToListNode(" [0,1,2]", -1),
				k:    4,
			},
			want: StrToListNode("[2,0,1]", -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
