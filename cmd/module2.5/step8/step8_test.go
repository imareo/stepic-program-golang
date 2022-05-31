package main

import (
	"reflect"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"около миши молоко", args{"около миши молоко"}, true},
		{"Около миши молоко", args{"Около миши молоко"}, true},
		{"Около Миши молоко", args{"Около Миши молоко"}, true},
		{"А роза упала на лапу Азора", args{"А роза упала на лапу Азора"}, true},
		{"Колбаса", args{"Колбаса"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
