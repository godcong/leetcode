package query

import (
	"testing"
)

func TestWriteMarkdownTo(t *testing.T) {
	type args struct {
		path string
		code *Code
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				path: "test.md",
				code: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, err := GetCode()
			if err != nil {
				t.Fatal(err)
			}

			if err := WriteMarkdownTo(tt.args.path, code); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
