package _1276

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_numOfBurgers(t *testing.T) {
	type args struct {
		tomatoSlices int
		cheeseSlices int
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
			if got := numOfBurgers(tt.args.tomatoSlices, tt.args.cheeseSlices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numOfBurgers() = %v, want %v", got, tt.want)
			}
		})
	}
}
