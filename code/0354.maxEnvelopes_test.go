package code

import "testing"

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				envelopes: strToIntArrArray(" [[5,4],[6,4],[6,7],[2,3]]"),
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				envelopes: strToIntArrArray("[[1,1],[1,1],[1,1]]"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
