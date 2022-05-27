/*
Заданы три числа - a,b,c(a<b<c) - длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным. Если является, вывести
"Прямоугольный". Иначе вывести "Непрямоугольный"
Sample Input:6 8 10
Sample Output:Прямоугольный
*/
package main

import "fmt"

func isRect(a, b, c int) string {
	if a*a+b*b == c*c {
		return "Прямоугольный"
	}
	return "Непрямоугольный"
}

func main() {
	var a, b, c int

	_, _ = fmt.Scan(&a, &b, &c)
	fmt.Print(isRect(a, b, c))
}
