package _1177

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canMakePaliQueries(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakePaliQueries(tt.args.s, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canMakePaliQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
