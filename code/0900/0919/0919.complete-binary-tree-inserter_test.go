package _0919

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
		want CBTInserter
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

func TestCBTInserter_Insert(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *CBTInserter
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Insert(tt.args.val); got != tt.want {
				t.Errorf("CBTInserter.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCBTInserter_Get_root(t *testing.T) {
	tests := []struct {
		name string
		this *CBTInserter
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get_root(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CBTInserter.Get_root() = %v, want %v", got, tt.want)
			}
		})
	}
}
