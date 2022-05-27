/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Sample Input:564 8954
Sample Output:5 4
*/
package main

import "fmt"

func main() {
	var num1, num2 string
	_, _ = fmt.Scan(&num1, &num2)
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			if num1[i] == num2[j] {
				fmt.Printf("%c ", num1[i])
			}
		}
	}
}
