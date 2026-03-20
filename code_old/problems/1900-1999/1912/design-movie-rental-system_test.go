package _1912

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n       int
		entries [][]int
	}
	tests := []struct {
		name string
		args args
		want MovieRentingSystem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.n, tt.args.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieRentingSystem_Search(t *testing.T) {
	type args struct {
		movie int
	}
	tests := []struct {
		name string
		this *MovieRentingSystem
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Search(tt.args.movie); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRentingSystem.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMovieRentingSystem_Rent(t *testing.T) {
	type args struct {
		shop  int
		movie int
	}
	tests := []struct {
		name string
		this *MovieRentingSystem
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Rent(tt.args.shop, tt.args.movie)
		})
	}
}

func TestMovieRentingSystem_Drop(t *testing.T) {
	type args struct {
		shop  int
		movie int
	}
	tests := []struct {
		name string
		this *MovieRentingSystem
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Drop(tt.args.shop, tt.args.movie)
		})
	}
}

func TestMovieRentingSystem_Report(t *testing.T) {
	tests := []struct {
		name string
		this *MovieRentingSystem
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Report(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovieRentingSystem.Report() = %v, want %v", got, tt.want)
			}
		})
	}
}
