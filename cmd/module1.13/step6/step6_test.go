package main

import "testing"

func Test_isExist(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"6,8,10", args{6, 8, 10}, "Существует"},
		{"3,4,5", args{3, 4, 5}, "Существует"},
		{"1,2,4", args{1, 2, 4}, "Не существует"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExist(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("isExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
