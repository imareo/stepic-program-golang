package main

import "testing"

func Test_battery_String(t *testing.T) {
	type fields struct {
		charge string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"1000010011", fields{"1000010011"}, "[      XXXX]"},
		{"1100000011", fields{"1100000011"}, "[      XXXX]"},
		{"1111111111", fields{"1111111111"}, "[XXXXXXXXXX]"},
		{"0000000000", fields{"0000000000"}, "[          ]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := battery{
				charge: tt.fields.charge,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
