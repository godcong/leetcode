package _0368

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
