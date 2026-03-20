package _1040

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numMovesStonesII(t *testing.T) {
	type args struct {
		stones []int
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
			if got := numMovesStonesII(tt.args.stones); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStonesII() = %v, want %v", got, tt.want)
			}
		})
	}
}
