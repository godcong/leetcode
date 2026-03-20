package _3285

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_stableMountains(t *testing.T) {
	type args struct {
		height    []int
		threshold int
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
			if got := stableMountains(tt.args.height, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stableMountains() = %v, want %v", got, tt.want)
			}
		})
	}
}
