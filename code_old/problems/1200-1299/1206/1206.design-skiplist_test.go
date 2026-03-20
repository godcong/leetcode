package _1206

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Skiplist
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

func TestSkiplist_Search(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		this *Skiplist
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Search(tt.args.target); got != tt.want {
				t.Errorf("Skiplist.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkiplist_Add(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *Skiplist
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Add(tt.args.num)
		})
	}
}

func TestSkiplist_Erase(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		this *Skiplist
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Erase(tt.args.num); got != tt.want {
				t.Errorf("Skiplist.Erase() = %v, want %v", got, tt.want)
			}
		})
	}
}
