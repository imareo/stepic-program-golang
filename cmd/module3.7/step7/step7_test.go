package main

import (
	"testing"
)

func Test_timeDiff(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t2 before t1", args{"13.03.2018 14:00:15,12.03.2018 14:00:15"}, "24h0m0s"},
		{"t2 after t1", args{"12.03.2018 14:00:15,13.03.2018 14:00:15"}, "24h0m0s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeDiff(tt.args.t); got != tt.want {
				t.Errorf("timeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
