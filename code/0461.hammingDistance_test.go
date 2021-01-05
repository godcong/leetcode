package code

import "testing"

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				x: 1,
				y: 3,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				x: 93,
				y: 73,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
