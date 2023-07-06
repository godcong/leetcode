package _2178

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumEvenSplit(t *testing.T) {
	type args struct {
		finalSum int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumEvenSplit(tt.args.finalSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumEvenSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
