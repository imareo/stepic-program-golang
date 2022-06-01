package main

import "testing"

func Test_maxDigit(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"12345", args{"12345"}, 5},
		{"12534", args{"12534"}, 5},
		{"111122", args{"111122"}, 2},
		{"11112211", args{"11112211"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDigit(tt.args.num); got != tt.want {
				t.Errorf("maxDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
