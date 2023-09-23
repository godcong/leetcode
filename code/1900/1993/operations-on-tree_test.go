package _1993

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		parent []int
	}
	tests := []struct {
		name string
		args args
		want LockingTree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLockingTree_Lock(t *testing.T) {
	type args struct {
		num  int
		user int
	}
	tests := []struct {
		name string
		this *LockingTree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Lock(tt.args.num, tt.args.user); got != tt.want {
				t.Errorf("LockingTree.Lock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLockingTree_Unlock(t *testing.T) {
	type args struct {
		num  int
		user int
	}
	tests := []struct {
		name string
		this *LockingTree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Unlock(tt.args.num, tt.args.user); got != tt.want {
				t.Errorf("LockingTree.Unlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLockingTree_Upgrade(t *testing.T) {
	type args struct {
		num  int
		user int
	}
	tests := []struct {
		name string
		this *LockingTree
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Upgrade(tt.args.num, tt.args.user); got != tt.want {
				t.Errorf("LockingTree.Upgrade() = %v, want %v", got, tt.want)
			}
		})
	}
}
