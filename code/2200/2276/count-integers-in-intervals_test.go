package _2276

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want CountIntervals
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

func TestCountIntervals_Add(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this *CountIntervals
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.left, tt.args.right)
		})
	}
}

func TestCountIntervals_Count(t *testing.T) {
	tests := []struct {
		name string
		this *CountIntervals
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Count(); got != tt.want {
				t.Errorf("CountIntervals.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
