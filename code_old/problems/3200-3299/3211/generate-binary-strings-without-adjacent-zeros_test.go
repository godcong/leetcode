package _3211

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_validStrings(t *testing.T) {
	type args struct {
		n int
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
			if got := validStrings(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
