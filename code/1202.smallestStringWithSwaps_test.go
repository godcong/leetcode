package code

import "testing"

func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				s:     "dcab",
				pairs: strToIntArrArray("[[0,3],[1,2]]"),
			},
			want: "bacd",
		},
		{
			name: "",
			args: args{
				s:     "dcab",
				pairs: strToIntArrArray("[[0,3],[1,2],[0,2]]"),
			},
			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
