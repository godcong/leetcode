package _0087

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
					t.Errorf("isScramble() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
