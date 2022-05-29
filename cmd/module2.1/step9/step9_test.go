package main

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{4}, 3},
		{"22", args{22}, 17711},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.num); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fibonacci(b *testing.B) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"22", args{22}, 17711},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			fibonacci(tt.args.num)
		}
	}
}
