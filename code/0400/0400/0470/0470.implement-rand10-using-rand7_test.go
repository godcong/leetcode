package _0470

import (
	"testing"
)

func Test_rand10(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rand10(); got != tt.want {
				t.Errorf("rand10() = %v, want %v", got, tt.want)
			}
		})
	}
}
