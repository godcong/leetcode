package _1851

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_minInterval(t *testing.T) {
	type args struct {
		intervals [][]int
		queries   []int
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
			if got := minInterval(tt.args.intervals, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
