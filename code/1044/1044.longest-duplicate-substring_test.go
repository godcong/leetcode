package _1044

import (
	"testing"
)

func Test_longestDupSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestDupSubstring(tt.args.s); got != tt.want {
				t.Errorf("longestDupSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
