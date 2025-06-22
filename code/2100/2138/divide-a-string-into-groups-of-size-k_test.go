package _2138

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideString() = %v, want %v", got, tt.want)
			}
		})
	}
}
