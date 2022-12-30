package _0855

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
		want ExamRoom
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

func TestExamRoom_Seat(t *testing.T) {
	tests := []struct {
		name string
		this *ExamRoom
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Seat(); got != tt.want {
				t.Errorf("ExamRoom.Seat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExamRoom_Leave(t *testing.T) {
	type args struct {
		p int
	}
	tests := []struct {
		name string
		this *ExamRoom
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Leave(tt.args.p)
		})
	}
}
