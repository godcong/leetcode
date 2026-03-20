package _1673

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_mostCompetitive(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := mostCompetitive(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostCompetitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
