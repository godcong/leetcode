package _1486

import "testing"

func Test_xorOperation(t *testing.T) {
	type args struct {
		n     int
		start int
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
				n:     5,
				start: 0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
					t.Errorf("xorOperation() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
