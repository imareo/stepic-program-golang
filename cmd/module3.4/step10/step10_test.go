package main

import (
	"fmt"
	"testing"
)

func Test_setValue(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"float64", args{1.2}, 1.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setValue(tt.args.val); got != tt.want {
				t.Errorf("setValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setValuePanic(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"panic int", args{1}},
		{"panic string", args{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("return with panic: %v\n", err)
				}
			}()
			setValue(tt.args.val)
		})
	}
}

func Test_setOperation(t *testing.T) {
	type args struct {
		op interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"+", args{"+"}, "+"},
		{"-", args{"-"}, "-"},
		{"*", args{"*"}, "*"},
		{"/", args{"/"}, "/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setOperation(tt.args.op); got != tt.want {
				t.Errorf("setOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setOperationPanic(t *testing.T) {
	type args struct {
		op interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"panic int", args{1}},
		{"panic float", args{1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("return with panic: %v\n", err)
				}
			}()
			setOperation(tt.args.op)
		})
	}
}

func TestDataset_getResult(t *testing.T) {
	type fields struct {
		value1    float64
		value2    float64
		operation string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"2.0 + 3.0", fields{2.0, 3.0, "+"}, 5.000},
		{"2.0 - 3.0", fields{2.0, 3.0, "-"}, -1},
		{"2.0 * 3.0", fields{2.0, 3.0, "*"}, 6},
		{"3.0 / 2.0", fields{3.0, 2.0, "/"}, 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dataset{
				value1:    tt.fields.value1,
				value2:    tt.fields.value2,
				operation: tt.fields.operation,
			}
			if got := d.getResult(); got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataset_getResultPanic(t *testing.T) {
	type fields struct {
		value1    float64
		value2    float64
		operation string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"panic 2.0 // 3.0", fields{2.0, 3.0, "//"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("return with panic: %v\n", err)
				}
			}()
			d := &Dataset{
				value1:    tt.fields.value1,
				value2:    tt.fields.value2,
				operation: tt.fields.operation,
			}
			d.getResult()
		})
	}
}
