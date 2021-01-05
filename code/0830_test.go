package code

import (
	"reflect"
	"testing"
)

func Test_largeGroupPositions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s: "abbxxxxzzy",
			},
			want: strToIntArrArray("[[3,6]]"),
		},
		{
			name: "",
			args: args{
				s: "abc",
			},
			want: strToIntArrArray("[]"),
		},
		{
			name: "",
			args: args{
				s: "abcdddeeeeaabbbcd",
			},
			want: strToIntArrArray("[[3,5],[6,9],[12,14]]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
