package _0707

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyLinkedList
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

func TestMyLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.index); got != tt.want {
				t.Errorf("MyLinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtHead(tt.args.val)
		})
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtTail(tt.args.val)
		})
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddAtIndex(tt.args.index, tt.args.val)
		})
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *MyLinkedList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.DeleteAtIndex(tt.args.index)
		})
	}
}
