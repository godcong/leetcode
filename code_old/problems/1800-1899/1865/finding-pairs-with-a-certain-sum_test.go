package _1865

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want FindSumPairs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSumPairs_Add(t *testing.T) {
	type args struct {
		index int
		val   int
	}
	tests := []struct {
		name string
		this *FindSumPairs
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.index, tt.args.val)
		})
	}
}

func TestFindSumPairs_Count(t *testing.T) {
	type args struct {
		tot int
	}
	tests := []struct {
		name string
		this *FindSumPairs
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Count(tt.args.tot); got != tt.want {
				t.Errorf("FindSumPairs.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
