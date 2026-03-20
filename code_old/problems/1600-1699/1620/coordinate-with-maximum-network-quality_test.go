package _1620

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_bestCoordinate(t *testing.T) {
	type args struct {
		towers [][]int
		radius int
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
			if got := bestCoordinate(tt.args.towers, tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bestCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}
