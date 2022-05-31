package main

import "testing"

func Test_getEven(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ihgewlqlkot->hello", args{"ihgewlqlkot"}, "hello"},
		{"ih->h", args{"ih"}, "h"},
		{"i-> ", args{"i"}, ""},
		{" -> ", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEven(tt.args.s); got != tt.want {
				t.Errorf("getEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
