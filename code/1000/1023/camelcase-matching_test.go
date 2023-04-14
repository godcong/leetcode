package _1023

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_camelMatch(t *testing.T) {
	type args struct {
		queries []string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := camelMatch(tt.args.queries, tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("camelMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
