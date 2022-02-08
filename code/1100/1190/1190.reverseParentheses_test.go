package _1190

import (
	"testing"
)

func Test_reverseParentheses(t *testing.T) {
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
			if got := reverseParentheses(tt.args.s); got != tt.want {
				t.Errorf("reverseParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
