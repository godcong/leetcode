package _2296

import (
	"reflect"
	"testing"

	. "github.com/godcong/leetcode/common"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want TextEditor
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

func TestTextEditor_AddText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		this *TextEditor
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.this.AddText(tt.args.text)
		})
	}
}

func TestTextEditor_DeleteText(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		this *TextEditor
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.DeleteText(tt.args.k); got != tt.want {
				t.Errorf("TextEditor.DeleteText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextEditor_CursorLeft(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		this *TextEditor
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.CursorLeft(tt.args.k); got != tt.want {
				t.Errorf("TextEditor.CursorLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextEditor_CursorRight(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		this *TextEditor
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.CursorRight(tt.args.k); got != tt.want {
				t.Errorf("TextEditor.CursorRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
