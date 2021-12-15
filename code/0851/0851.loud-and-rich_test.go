package _0851

import (
	"reflect"
	"testing"
)

func Test_loudAndRich(t *testing.T) {
	type args struct {
		richer [][]int
		quiet  []int
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
			if got := loudAndRich(tt.args.richer, tt.args.quiet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loudAndRich() = %v, want %v", got, tt.want)
			}
		})
	}
}
