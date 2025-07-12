package _1900

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_earliestAndLatest(t *testing.T) {
	type args struct {
		n            int
		firstPlayer  int
		secondPlayer int
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
			if got := earliestAndLatest(tt.args.n, tt.args.firstPlayer, tt.args.secondPlayer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("earliestAndLatest() = %v, want %v", got, tt.want)
			}
		})
	}
}
