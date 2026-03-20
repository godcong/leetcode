package _2266

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_countTexts(t *testing.T) {
	type args struct {
		pressedKeys string
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
			if got := countTexts(tt.args.pressedKeys); got != tt.want {
				t.Errorf("countTexts() = %v, want %v", got, tt.want)
			}
		})
	}
}
