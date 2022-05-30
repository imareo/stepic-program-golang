/*
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные
нужно вывести.

Sample Input:2 4
Sample Output:4 2
*/
package main

import "fmt"

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, *x2)
}
