package _1185

import (
	"testing"
)

func Test_dayOfTheWeek(t *testing.T) {
	type args struct {
		day   int
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfTheWeek(tt.args.day, tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
