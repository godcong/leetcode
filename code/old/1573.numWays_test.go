package old

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		s string
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
				s: "10101",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "1001",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s: "0000",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s: "100100010100110",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.s); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
