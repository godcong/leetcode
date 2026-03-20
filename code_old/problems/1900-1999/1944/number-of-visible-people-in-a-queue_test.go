package _1944

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_canSeePersonsCount(t *testing.T) {
	type args struct {
		heights []int
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
			if got := canSeePersonsCount(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canSeePersonsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
