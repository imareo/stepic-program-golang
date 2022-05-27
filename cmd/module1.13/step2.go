/*
Дано трехзначное число. Найдите сумму его цифр.
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:745
Sample Output:16
*/
package main

import "fmt"

func sumDigitStr(num string) int {
	res := 0
	for i := 0; i < len(num); i++ {
		res += int(num[i] - '0')
	}
	return res
}

func sumDigitInt(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	fmt.Println(sumDigitStr(fmt.Sprintf("%d", num)))
	// or fmt.Println(sumDigitInt(num))
}
