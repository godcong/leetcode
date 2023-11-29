package _2336

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want SmallestInfiniteSet
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

func TestSmallestInfiniteSet_PopSmallest(t *testing.T) {
	tests := []struct {
		name string
		this *SmallestInfiniteSet
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.PopSmallest(); got != tt.want {
				t.Errorf("SmallestInfiniteSet.PopSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallestInfiniteSet_AddBack(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *SmallestInfiniteSet
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddBack(tt.args.num)
		})
	}
}
