package main

import "testing"

func Test_sumDigitInt(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"123", args{123}, 6},
		{"745", args{745}, 16},
		{"111", args{111}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitInt(tt.args.num); got != tt.want {
				t.Errorf("sumDigitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumDigitStr(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"123", args{"123"}, 6},
		{"745", args{"745"}, 16},
		{"111", args{"111"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigitStr(tt.args.num); got != tt.want {
				t.Errorf("sumDigitStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_sumDigitStr(b *testing.B) {
	tests := []struct {
		args string
	}{
		{"123"},
		{"745"},
		{"111"},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			sumDigitStr(tt.args)
		}
	}
}

func Benchmark_sumDigitInt(b *testing.B) {
	tests := []struct {
		args int
	}{
		{123},
		{745},
		{111},
	}
	for _, tt := range tests {

		for i := 0; i < b.N; i++ {
			sumDigitInt(tt.args)
		}
	}
}
