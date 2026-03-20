package _2595

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_evenOddBit(t *testing.T) {
	type args struct {
		n int
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
			if got := evenOddBit(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
