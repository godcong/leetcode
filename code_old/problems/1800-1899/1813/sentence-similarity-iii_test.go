package _1813

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_areSentencesSimilar(t *testing.T) {
	type args struct {
		sentence1 string
		sentence2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.args.sentence1, tt.args.sentence2); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
