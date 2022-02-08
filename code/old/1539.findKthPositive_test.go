package old

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{
					2, 3, 4, 7, 11,
				},
				k: 5,
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				arr: []int{
					1, 2, 3, 4,
				},
				k: 2,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				arr: []int{
					1, 2,
				},
				k: 1,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				arr: []int{
					4, 5,
				},
				k: 1,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				arr: []int{
					4, 5,
				},
				k: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
