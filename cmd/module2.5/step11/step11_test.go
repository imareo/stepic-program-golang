package main

import "testing"

func Test_delChar(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"lcolmtnatl", args{"lcolmtnatl"}, "comna"},
		{"zaabcbd", args{"zaabcbd"}, "zcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delMultiChar(tt.args.str); got != tt.want {
				t.Errorf("delChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
