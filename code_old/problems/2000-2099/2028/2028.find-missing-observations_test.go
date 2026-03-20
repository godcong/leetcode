package _2028

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
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
			if got := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
