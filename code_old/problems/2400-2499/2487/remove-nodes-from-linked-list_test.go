package _2487

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_removeNodes(t *testing.T) {
	type args struct {
		head *ListNode
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
			if got := removeNodes(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
