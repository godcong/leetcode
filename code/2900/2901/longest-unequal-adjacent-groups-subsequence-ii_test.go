package _2901

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	type args struct {
		words  []string
		groups []int
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
			if got := getWordsInLongestSubsequence(tt.args.words, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getWordsInLongestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
