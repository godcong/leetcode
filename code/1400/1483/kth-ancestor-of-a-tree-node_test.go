package _1483

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n      int
		parent []int
	}
	tests := []struct {
		name string
		args args
		want TreeAncestor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeAncestor_GetKthAncestor(t *testing.T) {
	type args struct {
		node int
		k    int
	}
	tests := []struct {
		name string
		this *TreeAncestor
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.GetKthAncestor(tt.args.node, tt.args.k); got != tt.want {
				t.Errorf("TreeAncestor.GetKthAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
