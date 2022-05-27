package main

import "testing"

func Test_getHourMinute(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name  string
		args  args
		wantH int
		wantM int
	}{
		{"80000", args{80000}, 22, 13},
		{"13257", args{13257}, 3, 40},
		{"61", args{61}, 0, 1},
		{"57", args{57}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotM := getHourMinute(tt.args.k)
			if gotH != tt.wantH {
				t.Errorf("getHourMinute() gotH = %v, want %v", gotH, tt.wantH)
			}
			if gotM != tt.wantM {
				t.Errorf("getHourMinute() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}
