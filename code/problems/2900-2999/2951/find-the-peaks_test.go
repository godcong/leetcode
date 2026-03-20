package _2951

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_findPeaks(t *testing.T) {
	type args struct {
		mountain []int
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
			if got := findPeaks(tt.args.mountain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
