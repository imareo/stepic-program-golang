package main

import "testing"

func Test_toBinConv(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1}, "1"},
		{"2", args{2}, "10"},
		{"12", args{12}, "1100"},
		{"100", args{100}, "1100100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toBinConv(tt.args.num); got != tt.want {
				t.Errorf("toBinConv() = %v, want %v", got, tt.want)
			}
		})
	}
}
