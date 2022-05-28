package main

import "testing"

func Test_delNum(t *testing.T) {
	type args struct {
		num string
		del string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"38012732 1", args{"38012732", "1"}, "3802732"},
		{"38012732 2", args{"38012732", "2"}, "380173"},
		{"38012732 3", args{"38012732", "3"}, "801272"},
		{"38012732 4", args{"38012732", "4"}, "38012732"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNum(tt.args.num, tt.args.del); got != tt.want {
				t.Errorf("delNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
