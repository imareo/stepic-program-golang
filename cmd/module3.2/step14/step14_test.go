package main

import "testing"

func Test_csvTo2Float(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{"1 149,6088607594936;1 179,0666666666666", args{"1 149,6088607594936;1 179,0666666666666"},
			1149.6088607594936, 1179.0666666666666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := csvTo2Float(tt.args.s)
			if got != tt.want {
				t.Errorf("csvTo2Float() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("csvTo2Float() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
