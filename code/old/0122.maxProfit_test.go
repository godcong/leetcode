package old

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
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
				prices: strToIntArray("[7,1,5,3,6,4]"),
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				prices: strToIntArray("[1,2,3,4,5]"),
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				prices: strToIntArray("[7,6,4,3,1]"),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
