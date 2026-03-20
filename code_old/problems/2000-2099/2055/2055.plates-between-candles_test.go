package _2055

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_platesBetweenCandles(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
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
			if got := platesBetweenCandles(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("platesBetweenCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
