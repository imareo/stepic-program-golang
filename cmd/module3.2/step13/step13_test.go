package main

import "testing"

func Test_parseStrToInt1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"%^80", args{"%^80"}, 80},
		{"hhhhh20&&&&nd", args{"hhhhh20&&&&nd"}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseStrToInt(tt.args.s); got != tt.want {
				t.Errorf("parseStrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
