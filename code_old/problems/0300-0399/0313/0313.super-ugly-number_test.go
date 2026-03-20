package _0313

import (
	"testing"
)

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
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
			if got := nthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
