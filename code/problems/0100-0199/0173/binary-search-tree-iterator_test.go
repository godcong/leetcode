package _0173

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want BSTIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator_Next(t *testing.T) {
	tests := []struct {
		name string
		this *BSTIterator
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Next(); got != tt.want {
				t.Errorf("BSTIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator_HasNext(t *testing.T) {
	tests := []struct {
		name string
		this *BSTIterator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.HasNext(); got != tt.want {
				t.Errorf("BSTIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
