package _1053

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_prevPermOpt1(t *testing.T) {
	type args struct {
		arr []int
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
			if got := prevPermOpt1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevPermOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
