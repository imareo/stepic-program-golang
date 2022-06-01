/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат
деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка",
если ошибки нет - результат функции. Функция divide(a int, b int) (int, error) получает на вход
два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже
объявлен, а нужные пакеты уже импортированы!

Sample Input:10 5
Sample Output:2
*/
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func main() {
	var num1, num2 int

	_, _ = fmt.Scan(&num1, &num2)
	res, err := divide(num1, num2)
	if err != nil {
		fmt.Println(res)
	} else {
		fmt.Println("ошибка")
	}
}
