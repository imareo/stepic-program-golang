package main

import "testing"

func Test_isCorrectPass(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"fdsghdfgjsdDD1", args{"fdsghdfgjsdDD1"}, true},
		{"fdsgh", args{"fdsgh"}, true},
		{"fdsg", args{"fdsg"}, false},
		{"fdsg1", args{"fdsg1"}, true},
		{"fdsg+", args{"fdsg+"}, false},
		{"fds g", args{"fds g"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCorrectPass(tt.args.str); got != tt.want {
				t.Errorf("isCorrectPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
