package main

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"123", args{"123"}, "321"},
		{"745", args{"745"}, "547"},
		{"111", args{"111"}, "111"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.num); got != tt.want {
				t.Errorf("sumDigitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
