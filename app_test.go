package main

import (
	"os"
	"testing"
)

func Test_execScript(t *testing.T) {
	pwd := os.Getenv("PWD")
	tests := []struct {
		name     string
		path     string
		wantCode int
		wantOut  string
		wantErr  string
	}{
		{"sanity", pwd + "/scripts_template/testing_script.sh", 0, "This is a test\n", ""},
		{"exit1", pwd + "/scripts_template/exit_1.sh", 1, "", ""},
		{"stderr", pwd + "/scripts_template/stderr.sh", 0, "", "test\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt)
			gotCode, gotOut, gotErr, err := execScript(tt.path)
			if gotCode != tt.wantCode {
				t.Errorf("execScript() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotOut != tt.wantOut {
				t.Errorf("execScript() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotErr != tt.wantErr {
				t.Errorf("execScript() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			if err != nil {
				t.Errorf("execScript() err = %v, want <nil>", err)
			}
		})
	}
}
