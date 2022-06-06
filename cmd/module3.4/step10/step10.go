package main

import (
	"fmt"
)

type Dataset struct {
	value1    float64
	value2    float64
	operation string
}

func readTask() (x, y, op interface{}) {
	return 1.0, 2.0, "/"
}

func setOperation(op interface{}) string {
	res, ok := op.(string)
	if !ok {
		panic(fmt.Sprintf("value=%v: %T", op, op))
	}
	return res
}

func setValue(val interface{}) float64 {
	res, ok := val.(float64)
	if !ok {
		panic(fmt.Sprintf("value=%v: %T", val, val))
	}
	return res
}

func (d *Dataset) getResult() float64 {
	switch d.operation {
	case "+":
		return d.value1 + d.value2
	case "-":
		return d.value1 - d.value2
	case "*":
		return d.value1 * d.value2
	case "/":
		return d.value1 / d.value2
	}
	panic("неизвестная операция")
}

func main() {
	value1, value2, operation := readTask()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ds := Dataset{
		setValue(value1),
		setValue(value2),
		setOperation(operation),
	}
	fmt.Printf("%.4f", ds.getResult())
}
