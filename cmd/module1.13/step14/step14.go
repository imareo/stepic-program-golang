/*
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи,
выведите число -1.

Sample Input:8
Sample Output:6
*/
package main

import "fmt"

func fib(num int) int {
	a, b := 1, 1
	for i := 3; b <= num; i++ {
		if b == num {
			return i - 1
		}
		a, b = b, a+b
	}
	return -1
}

func main() {
	var num int

	_, _ = fmt.Scan(&num)
	fmt.Println(fib(num))
}
