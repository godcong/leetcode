package _1846

import (
	"testing"
)

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumElementAfterDecrementingAndRearranging(tt.args.arr); got != tt.want {
				t.Errorf("maximumElementAfterDecrementingAndRearranging() = %v, want %v", got, tt.want)
			}
		})
	}
}
