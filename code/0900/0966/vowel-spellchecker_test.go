package _0966

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_spellchecker(t *testing.T) {
	type args struct {
		wordlist []string
		queries  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spellchecker(tt.args.wordlist, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spellchecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
