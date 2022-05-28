package main

import "testing"

func Test_getNumeral(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{1}, "1 korova"},
		{"2", args{2}, "2 korovy"},
		{"11", args{11}, "11 korov"},
		{"12", args{12}, "12 korov"},
		{"22", args{22}, "22 korovy"},
		{"30", args{30}, "30 korov"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumeral(tt.args.num); got != tt.want {
				t.Errorf("getNumeral() = %v, want %v", got, tt.want)
			}
		})
	}
}
