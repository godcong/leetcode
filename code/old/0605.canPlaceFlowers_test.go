package old

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				flowerbed: []int{
					1, 0, 0, 0, 1,
				},
				n: 1,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{
					1, 0, 0, 0, 1,
				},
				n: 2,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{
					1, 0, 0, 0, 1, 0, 0,
				},
				n: 2,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				flowerbed: []int{
					0, 0, 0, 0, 0, 1, 0, 0,
				},
				n: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
