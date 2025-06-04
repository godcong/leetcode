package _3403

import (
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_answerString(t *testing.T) {
	type args struct {
		word       string
		numFriends int
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
			if got := answerString(tt.args.word, tt.args.numFriends); got != tt.want {
				t.Errorf("answerString() = %v, want %v", got, tt.want)
			}
		})
	}
}
