package _2807

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_insertGreatestCommonDivisors(t *testing.T) {
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
			if got := insertGreatestCommonDivisors(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertGreatestCommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
