package code

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		A []int
		K int
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
				A: strToIntArray("[0,1,0]"),
				K: 1,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,1,0]"),
				K: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
