package _0689

import (
	"reflect"
	"testing"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := maxSumOfThreeSubarrays(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSumOfThreeSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
