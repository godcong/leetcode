package _0622

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want MyCircularQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_EnQueue(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		this *MyCircularQueue
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.EnQueue(tt.args.value); got != tt.want {
				t.Errorf("MyCircularQueue.EnQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.DeQueue(); got != tt.want {
				t.Errorf("MyCircularQueue.DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Front(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Front(); got != tt.want {
				t.Errorf("MyCircularQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Rear(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Rear(); got != tt.want {
				t.Errorf("MyCircularQueue.Rear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsFull(t *testing.T) {
	tests := []struct {
		name string
		this *MyCircularQueue
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.IsFull(); got != tt.want {
				t.Errorf("MyCircularQueue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
