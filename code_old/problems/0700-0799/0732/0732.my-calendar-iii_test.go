package _0732

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyCalendarThree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalendarThree_Book(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		this *MyCalendarThree
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendarThree.Book() = %v, want %v", got, tt.want)
			}
		})
	}
}
