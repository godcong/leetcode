package _1670

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want FrontMiddleBackQueue
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

func TestFrontMiddleBackQueue_PushFront(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.PushFront(tt.args.val)
		})
	}
}

func TestFrontMiddleBackQueue_PushMiddle(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.PushMiddle(tt.args.val)
		})
	}
}

func TestFrontMiddleBackQueue_PushBack(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.PushBack(tt.args.val)
		})
	}
}

func TestFrontMiddleBackQueue_PopFront(t *testing.T) {
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.PopFront(); got != tt.want {
				t.Errorf("FrontMiddleBackQueue.PopFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrontMiddleBackQueue_PopMiddle(t *testing.T) {
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.PopMiddle(); got != tt.want {
				t.Errorf("FrontMiddleBackQueue.PopMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrontMiddleBackQueue_PopBack(t *testing.T) {
	tests := []struct {
		name string
		this *FrontMiddleBackQueue
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.PopBack(); got != tt.want {
				t.Errorf("FrontMiddleBackQueue.PopBack() = %v, want %v", got, tt.want)
			}
		})
	}
}
