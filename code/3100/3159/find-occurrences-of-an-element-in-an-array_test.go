package _3159

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_occurrencesOfElement(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
		x       int
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
			if got := occurrencesOfElement(tt.args.nums, tt.args.queries, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("occurrencesOfElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
