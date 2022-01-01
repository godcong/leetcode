package _2022

import (
	"reflect"
	"testing"
)

func Test_construct2DArray(t *testing.T) {
	type args struct {
		original []int
		m        int
		n        int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct2DArray(tt.args.original, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
