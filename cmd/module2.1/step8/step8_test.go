package main

import "testing"

func Test_vote(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0 0", args{0, 0, 0}, 0},
		{"0 0 1", args{0, 0, 1}, 0},
		{"0 1 1", args{0, 1, 1}, 1},
		{"1 1 1", args{1, 1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vote(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("vote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vote1(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0 0", args{0, 0, 0}, 0},
		{"0 0 1", args{0, 0, 1}, 0},
		{"0 1 1", args{0, 1, 1}, 1},
		{"1 1 1", args{1, 1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vote1(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("vote1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_vote(b *testing.B) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0 0", args{0, 0, 0}, 0},
		{"0 0 1", args{0, 0, 1}, 0},
		{"0 1 1", args{0, 1, 1}, 1},
		{"1 1 1", args{1, 1, 1}, 1},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			vote(tt.args.x, tt.args.y, tt.args.z)
		}
	}
}

func Benchmark_vote1(b *testing.B) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0 0", args{0, 0, 0}, 0},
		{"0 0 1", args{0, 0, 1}, 0},
		{"0 1 1", args{0, 1, 1}, 1},
		{"1 1 1", args{1, 1, 1}, 1},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			vote1(tt.args.x, tt.args.y, tt.args.z)
		}
	}
}
