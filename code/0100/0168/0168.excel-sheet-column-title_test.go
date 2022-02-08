package _0168

import (
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
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
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
