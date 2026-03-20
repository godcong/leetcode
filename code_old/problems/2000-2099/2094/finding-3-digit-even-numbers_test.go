package _2094

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findEvenNumbers(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findEvenNumbers(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findEvenNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
