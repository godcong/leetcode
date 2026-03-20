package _1765

import (
	"reflect"
	"testing"
)

func Test_highestPeak(t *testing.T) {
	type args struct {
		isWater [][]int
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
			if got := highestPeak(tt.args.isWater); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
