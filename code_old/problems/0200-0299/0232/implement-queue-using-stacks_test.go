package _0232

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyQueue
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

func TestMyQueue_Push(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		this *MyQueue
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

func TestMyQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *MyQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Pop(); got != tt.want {
				t.Errorf("MyQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	tests := []struct {
		name string
		this *MyQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Peek(); got != tt.want {
				t.Errorf("MyQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Empty(t *testing.T) {
	tests := []struct {
		name string
		this *MyQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Empty(); got != tt.want {
				t.Errorf("MyQueue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}
