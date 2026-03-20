package _1604

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_alertNames(t *testing.T) {
	type args struct {
		keyName []string
		keyTime []string
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
			if got := alertNames(tt.args.keyName, tt.args.keyTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alertNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
