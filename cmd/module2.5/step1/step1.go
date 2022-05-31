package main

import "fmt"

func main() {
	ExampleString()
}

func ExampleString() {
	var strR = "Это строка"
	var strE = "This is string"

	// срез строки работает по байтам
	fmt.Printf("Длины строк: %d байт и %d байт\n", len(strR), len(strE))
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\" и \"%v\"\n", strR[7:], strE[8:])

	for i, ch := range strR {
		fmt.Printf("[%v %c] ", i, ch)
	}
	fmt.Print("\n")

	for i, ch := range strE {
		fmt.Printf("[%v %c] ", i, ch)
	}
	fmt.Print("\n")

}
