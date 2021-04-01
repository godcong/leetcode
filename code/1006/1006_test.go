package _1006

import "testing"

func Test_clumsy(t *testing.T) {
	type args struct {
		N int
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
				N: 4,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.args.N); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
