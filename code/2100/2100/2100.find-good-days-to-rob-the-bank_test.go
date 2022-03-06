package _2100

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_goodDaysToRobBank(t *testing.T) {
	type args struct {
		security []int
		time     int
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
			if got := goodDaysToRobBank(tt.args.security, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("goodDaysToRobBank() = %v, want %v", got, tt.want)
			}
		})
	}
}
