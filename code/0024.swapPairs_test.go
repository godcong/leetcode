package code

import (
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
		deep int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: strToListNode("[1,2,3,4]", -1),
			},
			want: strToListNode("[2,1,4,3]", -1),
			deep: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !listNodeDeepEqual(t, got, tt.want, tt.deep) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
