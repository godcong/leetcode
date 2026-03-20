package _2682

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_circularGameLosers(t *testing.T) {
	type args struct {
		n int
		k int
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
			if got := circularGameLosers(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularGameLosers() = %v, want %v", got, tt.want)
			}
		})
	}
}
