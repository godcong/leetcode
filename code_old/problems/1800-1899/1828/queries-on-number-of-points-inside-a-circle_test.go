package _1828

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countPoints(t *testing.T) {
	type args struct {
		points  [][]int
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
			if got := countPoints(tt.args.points, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
