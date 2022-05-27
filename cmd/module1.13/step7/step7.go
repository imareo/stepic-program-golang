/*
Даны два числа. Найти их среднее арифметическое.
Sample Input 1:3 5
Sample Output 1:4
Sample Input 2:277 154
Sample Output 2:215.5
*/
package main

import "fmt"

func main() {
	var a, b float64

	_, _ = fmt.Scan(&a, &b)
	fmt.Print(average(a, b))
}

func average(a, b float64) float64 {
	return (a + b) / 2
}
