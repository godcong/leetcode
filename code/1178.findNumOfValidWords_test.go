package code

import (
	"reflect"
	"testing"
)

func Test_findNumOfValidWords(t *testing.T) {
	type args struct {
		words   []string
		puzzles []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				words:   strToStringArray("[\"aaaa\",\"asas\",\"able\",\"ability\",\"actt\",\"actor\",\"access\"]"),
				puzzles: strToStringArray("[\"aboveyz\",\"abrodyz\",\"abslute\",\"absoryz\",\"actresz\",\"gaswxyz\"]"),
			},
			want: strToIntArray("[1,1,3,2,4,0]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumOfValidWords(tt.args.words, tt.args.puzzles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumOfValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
