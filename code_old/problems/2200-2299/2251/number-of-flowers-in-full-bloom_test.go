package _2251

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_fullBloomFlowers(t *testing.T) {
	type args struct {
		flowers [][]int
		people  []int
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
			if got := fullBloomFlowers(tt.args.flowers, tt.args.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullBloomFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
