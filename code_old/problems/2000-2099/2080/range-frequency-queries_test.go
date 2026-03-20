package _2080

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want RangeFreqQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeFreqQuery_Query(t *testing.T) {
	type args struct {
		left  int
		right int
		value int
	}
	tests := []struct {
		name string
		this *RangeFreqQuery
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Query(tt.args.left, tt.args.right, tt.args.value); got != tt.want {
				t.Errorf("RangeFreqQuery.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
