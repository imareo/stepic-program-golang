/*
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует".

Sample Input:4 5 6
Sample Output:Существует
*/
package main

import "fmt"

func isExist(a, b, c int) string {
	if a+b < c || a+c < b || b+c < a {
		return "Не существует"
	}
	return "Существует"
}

func main() {
	var a, b, c int

	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Print(isExist(a, b, c))
}
