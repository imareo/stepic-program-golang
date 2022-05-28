/*
Найдите количество минимальных элементов в последовательности.
Вводится натуральное число N, а затем N целых чисел последовательности.
Sample Input:3
21 11 4
Sample Output:1
*/
package main

import (
	"fmt"
)

func main() {
	var num, temp, min, res int
	isFirst := true

	for _, _ = fmt.Scan(&num); num > 0; num-- {
		switch _, _ = fmt.Scan(&temp); {
		case isFirst:
			min, res, isFirst = temp, 1, false
		case temp == min:
			res++
		case temp < min:
			min, res = temp, 1
		}
	}
	fmt.Println(res)
}
