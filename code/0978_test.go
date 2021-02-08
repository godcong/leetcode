package code

import "testing"

func Test_maxTurbulenceSize(t *testing.T) {
	type args struct {
		arr []int
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
				arr: strToIntArray("[9,4,2,10,7,8,8,1,9]"),
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				arr: strToIntArray("[4,8,12,16]"),
			},
			want: 2,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSize(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
