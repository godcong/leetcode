package _1054

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_rearrangeBarcodes(t *testing.T) {
	type args struct {
		barcodes []int
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
			if got := rearrangeBarcodes(tt.args.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
