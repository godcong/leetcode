package _1093

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_sampleStats(t *testing.T) {
	type args struct {
		count []int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sampleStats(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sampleStats() = %v, want %v", got, tt.want)
			}
		})
	}
}
