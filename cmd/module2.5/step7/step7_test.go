package main

import "testing"

func Test_solver(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Быть или не быть.", args{"Быть или не быть."}, "Right"},
		{"быть или не быть.", args{"быть или не быть."}, "Wrong"},
		{"Быть или не быть", args{"Быть или не быть"}, "Wrong"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solver(tt.args.text); got != tt.want {
				t.Errorf("solver() = %v, want %v", got, tt.want)
			}
		})
	}
}
