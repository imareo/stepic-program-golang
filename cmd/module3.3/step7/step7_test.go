package main

import "testing"

func Test_forTest(t *testing.T) {
	type args struct {
		num uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"727178", args{727178}, 28},
		{"727170", args{727170}, 2},
		{"222", args{222}, 222},
		{"777", args{777}, 100},
		{"0", args{0}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forTest(tt.args.num); got != tt.want {
				t.Errorf("forTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
