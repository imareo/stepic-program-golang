package main

import (
	"errors"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{"10 5", args{10, 5}, 2, nil},
		{"10 0", args{10, 0}, 0, errors.New("divide by zero")},
		{"0 5", args{0, 5}, 0, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := divide(tt.args.a, tt.args.b)
			if (err != nil) && errors.Is(errors.New("divide by zero"), tt.wantErr) {
				t.Errorf("divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("divide() got = %v, want %v", got, tt.want)
			}
		})
	}
}
