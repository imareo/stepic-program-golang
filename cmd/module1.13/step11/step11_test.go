package main

import "testing"

func Test_solver(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"100 500", args{100, 500}, 497, true},
		{"-100 -92", args{-100, -92}, -98, true},
		{"-1 6", args{-1, 6}, 0, true},
		{"11 13", args{11, 13}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := solver(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("solver() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("solver() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
