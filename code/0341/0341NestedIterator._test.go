package _0341

import (
	. "github.com/godcong/leetcode/common"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nestedList []*NestedInteger
	}
	tests := []struct {
		name string
		args args
		want *NestedIterator
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nestedList: nil,
			},
			want: &NestedIterator{
				vals: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nestedList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
