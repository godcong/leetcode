package _1669

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
