package _2349

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want NumberContainers
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

func TestNumberContainers_Change(t *testing.T) {
	type args struct {
		index  int
		number int
	}
	tests := []struct {
		name string
		this *NumberContainers
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Change(tt.args.index, tt.args.number)
		})
	}
}

func TestNumberContainers_Find(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		this *NumberContainers
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Find(tt.args.number); got != tt.want {
				t.Errorf("NumberContainers.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
