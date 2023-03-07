package _1096

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_braceExpansionII(t *testing.T) {
	type args struct {
		expression string
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
			if got := braceExpansionII(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII() = %v, want %v", got, tt.want)
			}
		})
	}
}
