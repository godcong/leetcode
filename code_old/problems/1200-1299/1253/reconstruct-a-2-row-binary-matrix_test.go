package _1253

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_reconstructMatrix(t *testing.T) {
	type args struct {
		upper  int
		lower  int
		colsum []int
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
			if got := reconstructMatrix(tt.args.upper, tt.args.lower, tt.args.colsum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
