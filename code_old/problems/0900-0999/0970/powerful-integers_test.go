package _0970

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
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
			if got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
