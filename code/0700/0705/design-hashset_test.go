package _0705

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyHashSet
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

func TestMyHashSet_Add(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *MyHashSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.key)
		})
	}
}

func TestMyHashSet_Remove(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *MyHashSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Remove(tt.args.key)
		})
	}
}

func TestMyHashSet_Contains(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *MyHashSet
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Contains(tt.args.key); got != tt.want {
				t.Errorf("MyHashSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
