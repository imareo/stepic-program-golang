package main

import "testing"

func Test_sumInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1 2 3 4 5 6 7", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 7, 28},
		{"1 2 6 ", args{[]int{1, 2, 6}}, 3, 9},
		{"0 0 1 ", args{[]int{0, 0, 1}}, 3, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := sumInt(tt.args.a...)
			if got != tt.want {
				t.Errorf("sumInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("sumInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
