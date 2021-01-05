package code

import (
	"reflect"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "catsanddog",
				wordDict: []string{
					"cat", "cats", "and", "sand", "dog",
				},
			},
			want: []string{
				"cat sand dog",
				"cats and dog",
			},
		},
		{
			name: "",
			args: args{
				s: "pineapplepenapple",
				wordDict: []string{
					"apple", "pen", "applepen", "pine", "pineapple",
				},
			},
			want: []string{
				"pine apple pen apple",
				"pine applepen apple",
				"pineapple pen apple",
			},
		},
		{
			name: "",
			args: args{
				s: "catsandog",
				wordDict: []string{
					"cats", "dog", "sand", "and", "cat",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
