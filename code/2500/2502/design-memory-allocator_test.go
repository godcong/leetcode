package _2502

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want Allocator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocator_Allocate(t *testing.T) {
	type args struct {
		size int
		mID  int
	}
	tests := []struct {
		name string
		this *Allocator
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Allocate(tt.args.size, tt.args.mID); got != tt.want {
				t.Errorf("Allocator.Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllocator_FreeMemory(t *testing.T) {
	type args struct {
		mID int
	}
	tests := []struct {
		name string
		this *Allocator
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.FreeMemory(tt.args.mID); got != tt.want {
				t.Errorf("Allocator.FreeMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
