package query

import (
	"testing"
)

func TestGetGroupName(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				num: "0",
			},
			want: "0000",
		},
		{
			name: "",
			args: args{
				num: "100",
			},
			want: "0100",
		},
		{
			name: "",
			args: args{
				num: "1000",
			},
			want: "1000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGroupName(tt.args.num); got != tt.want {
				t.Errorf("GetGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGroupName1(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				num: "剑指 Offer II 115",
			},
			want: "SwordRefers_Offer_II_0100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGroupName(tt.args.num); got != tt.want {
				t.Errorf("GetGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}
