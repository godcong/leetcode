package _0502

import (
	"testing"
)

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
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
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
