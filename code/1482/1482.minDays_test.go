package _1482

import "testing"

func Test_minDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
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
				bloomDay: nil,
				m:        0,
				k:        0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := minDays(tt.args.bloomDay, tt.args.m, tt.args.k); got != tt.want {
					t.Errorf("minDays() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
