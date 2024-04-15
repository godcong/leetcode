package _0706

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyHashMap
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

func TestMyHashMap_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		this *MyHashMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestMyHashMap_Get(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *MyHashMap
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Get(tt.args.key); got != tt.want {
				t.Errorf("MyHashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyHashMap_Remove(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		this *MyHashMap
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
