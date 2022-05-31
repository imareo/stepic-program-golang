/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского
алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok",
иначе вывести "Wrong password"

Sample Input:fdsghdfgjsdDD1
Sample Output:Ok
*/
package main

import (
	"fmt"
	"regexp"
)

func isCorrectPass(str string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z\d]{5,}$`, str)
	return matched
}

func main() {
	var str string

	_, _ = fmt.Scan(&str)
	if isCorrectPass(str) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
