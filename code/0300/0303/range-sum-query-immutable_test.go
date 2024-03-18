package _0303

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want NumArray
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this *NumArray
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
