package _1845

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want SeatManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeatManager_Reserve(t *testing.T) {
	tests := []struct {
		name string
		this *SeatManager
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Reserve(); got != tt.want {
				t.Errorf("SeatManager.Reserve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeatManager_Unreserve(t *testing.T) {
	type args struct {
		seatNumber int
	}
	tests := []struct {
		name string
		this *SeatManager
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Unreserve(tt.args.seatNumber)
		})
	}
}
