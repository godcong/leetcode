package code

import "testing"

func Test_countBinarySubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "00110011",
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s: "10101",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "0001110100010",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBinarySubstrings(tt.args.s); got != tt.want {
				t.Errorf("countBinarySubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
