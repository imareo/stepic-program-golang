package main

import "testing"

func Test_digRoot(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3456", args{3456}, 9},
		{"65536", args{65536}, 7},
		{"1", args{1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digRoot(tt.args.num); got != tt.want {
				t.Errorf("digRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
