package code

import (
	"testing"
)

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		A []int
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
				A: strToIntArray("[2,1,2]"),
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[1,2,1]"),
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[3,2,3,4]"),
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				A: strToIntArray("[3,6,2,3]"),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.A); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
