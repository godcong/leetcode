package _2766

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_relocateMarbles(t *testing.T) {
	type args struct {
		nums     []int
		moveFrom []int
		moveTo   []int
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
			if got := relocateMarbles(tt.args.nums, tt.args.moveFrom, tt.args.moveTo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relocateMarbles() = %v, want %v", got, tt.want)
			}
		})
	}
}
