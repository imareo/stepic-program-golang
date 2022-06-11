package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_timeCalc(t *testing.T) {
	type args struct {
		min time.Duration
		sec time.Duration
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"12 мин. 13 сек.", args{12, 13}, "Fri May 15 19:28:18 UTC 2020"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeCalc(tt.args.min, tt.args.sec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("timeCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
