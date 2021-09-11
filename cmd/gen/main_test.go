package main

import (
	"testing"
)

func Test_addToGit(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			args:    args{
				path: "E:\\workspace\\project\\leetcode\\code\\0600",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addToGit(tt.args.path,"add xxx"); (err != nil) != tt.wantErr {
				t.Errorf("addToGit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
