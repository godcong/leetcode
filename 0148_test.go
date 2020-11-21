package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
				head: strToListNode("[4,2,1,3]", -1),
			},
			want: strToListNode("[1,2,3,4]", -1),
		},
		{
			name: "",
			args: args{
				head: strToListNode("[-1,5,3,4,0]", -1),
			},
			want: strToListNode("[-1,0,3,4,5]", -1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
