package _0225

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Push(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		this *MyStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Push(tt.args.x)
		})
	}
}

func TestMyStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *MyStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Pop(); got != tt.want {
				t.Errorf("MyStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Top(t *testing.T) {
	tests := []struct {
		name string
		this *MyStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Top(); got != tt.want {
				t.Errorf("MyStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStack_Empty(t *testing.T) {
	tests := []struct {
		name string
		this *MyStack
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Empty(); got != tt.want {
				t.Errorf("MyStack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
