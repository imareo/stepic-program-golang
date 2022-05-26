/*
Дано натуральное число, выведите его последнюю цифру.
Формат входных данных
На вход дается натуральное число N, не превосходящее 10000.
Формат выходных данных
Выведите одно целое число - ответ на задачу.

Sample Input:123
Sample Output:3
*/
package main

import "fmt"

func main() {
	var num string

	_, _ = fmt.Scan(&num)
	res := num[len(num)-1]
	fmt.Println(string(res))
}
