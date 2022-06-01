package main

import "testing"

func Test_squaredDigit(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1"}, "1"},
		{"11", args{"11"}, "11"},
		{"12", args{"12"}, "14"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squaredDigit(tt.args.num); got != tt.want {
				t.Errorf("squaredDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
