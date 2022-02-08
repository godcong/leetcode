package old

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
		{
			name: "",
			args: args{
				5, 0,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				4, 3,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				1, 7,
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				10, 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
