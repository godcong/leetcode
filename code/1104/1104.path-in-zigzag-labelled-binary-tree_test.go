package _1104

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	type args struct {
		label int
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
			if got := pathInZigZagTree(tt.args.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
