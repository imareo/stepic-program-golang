package main

import "testing"

func Test_solveTime(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"before 13", args{"2020-05-15 08:00:00"}, "2020-05-15 08:00:00"},
		{"after  13", args{"2020-05-15 18:00:00"}, "2020-05-16 18:00:00"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveTime(tt.args.t); got != tt.want {
				t.Errorf("solveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
