package _2517

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_maximumTastiness(t *testing.T) {
	type args struct {
		price []int
		k     int
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
			if got := maximumTastiness(tt.args.price, tt.args.k); got != tt.want {
				t.Errorf("maximumTastiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
