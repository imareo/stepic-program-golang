/*
Дано натуральное число N. Выведите его представление в двоичном виде.

Sample Input:12
Sample Output:1100
*/
package main

import "fmt"

func toBinConv(num int) string {
	return fmt.Sprintf("%b", num)
}

func main() {
	var num int

	_, _ = fmt.Scan(&num)
	fmt.Println(toBinConv(num))
}
