package _1154

import (
	"testing"
)

func Test_dayOfYear(t *testing.T) {
	type args struct {
		date string
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
				date: "2019-01-09",
			},
			want: 9,
		},
		{
			name: "",
			args: args{
				date: "2019-02-10",
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfYear(tt.args.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
