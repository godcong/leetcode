package _1001

import (
	"reflect"
	"testing"
)

func Test_gridIllumination(t *testing.T) {
	type args struct {
		n       int
		lamps   [][]int
		queries [][]int
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
			if got := gridIllumination(tt.args.n, tt.args.lamps, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridIllumination() = %v, want %v", got, tt.want)
			}
		})
	}
}
