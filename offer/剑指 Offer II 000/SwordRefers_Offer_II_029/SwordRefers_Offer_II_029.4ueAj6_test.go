package _SwordRefers_Offer_II_029

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_insert(t *testing.T) {
	type args struct {
		aNode *Node
		x     int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.aNode, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
