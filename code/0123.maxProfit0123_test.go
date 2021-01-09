package code

import "testing"

func Test_maxProfit0123(t *testing.T) {
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
				prices: strToIntArray("[3,3,5,0,0,3,1,4]"),
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				prices: strToIntArray("[1,2,3,4,5]"),
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit0123(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit0123() = %v, want %v", got, tt.want)
			}
		})
	}
}
