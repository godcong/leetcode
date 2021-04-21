package _0363

import "testing"

func Test_maxInt(t *testing.T) {
	type args struct {
		a int
		b int
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
				a: 0,
				b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := maxInt(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("maxInt() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_maxSumSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := maxSumSubmatrix(tt.args.matrix, tt.args.k); got != tt.want {
					t.Errorf("maxSumSubmatrix() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_subMax(t *testing.T) {
	type args struct {
		sumArry []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := subMax(tt.args.sumArry, tt.args.k); got != tt.want {
					t.Errorf("subMax() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
