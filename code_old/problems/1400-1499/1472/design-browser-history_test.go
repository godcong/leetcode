package _1472

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	type args struct {
		homepage string
	}
	tests := []struct {
		name string
		args args
		want BrowserHistory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.homepage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Visit(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.Visit(tt.args.url)
		})
	}
}

func TestBrowserHistory_Back(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Back(tt.args.steps); got != tt.want {
				t.Errorf("BrowserHistory.Back() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrowserHistory_Forward(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		this *BrowserHistory
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Forward(tt.args.steps); got != tt.want {
				t.Errorf("BrowserHistory.Forward() = %v, want %v", got, tt.want)
			}
		})
	}
}
