package _0190

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				num: 0b00000010100101000001111010011100,
			},
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "",
			args: args{
				num: 0b11111111111111111111111111111101,
			},
			want: 0b10111111111111111111111111111111,
		},
		{
			name: "",
			args: args{
				num: 0b00000010100101000001111010011100,
			},
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "",
			args: args{
				num: 0b11111111111111111111111111111101,
			},
			want: 0b10111111111111111111111111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
