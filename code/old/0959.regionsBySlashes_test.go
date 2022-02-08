package old

import "testing"

func Test_regionsBySlashes(t *testing.T) {
	type args struct {
		grid []string
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
				grid: []string{
					" /",
					"/ ",
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				grid: []string{
					" /",
					"  ",
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				grid: []string{
					"\\/",
					"/\\",
				},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				grid: []string{
					"/\\",
					"\\/",
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				grid: []string{
					"//",
					"/ ",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regionsBySlashes(tt.args.grid); got != tt.want {
				t.Errorf("regionsBySlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
