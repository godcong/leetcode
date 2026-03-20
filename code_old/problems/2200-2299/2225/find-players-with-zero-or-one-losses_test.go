package _2225

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findWinners(t *testing.T) {
	type args struct {
		matches [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWinners(tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWinners() = %v, want %v", got, tt.want)
			}
		})
	}
}
