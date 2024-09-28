package _2286

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want BookMyShow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookMyShow_Gather(t *testing.T) {
	type args struct {
		k      int
		maxRow int
	}
	tests := []struct {
		name string
		this *BookMyShow
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Gather(tt.args.k, tt.args.maxRow); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BookMyShow.Gather() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBookMyShow_Scatter(t *testing.T) {
	type args struct {
		k      int
		maxRow int
	}
	tests := []struct {
		name string
		this *BookMyShow
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Scatter(tt.args.k, tt.args.maxRow); got != tt.want {
				t.Errorf("BookMyShow.Scatter() = %v, want %v", got, tt.want)
			}
		})
	}
}
