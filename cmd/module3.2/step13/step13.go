/*
Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал
модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию
которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением
этих строк.
Вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки.
Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы.
Ваша функция должна называться adding() !
Функция и пакет main уже объявлены!

Sample Input:%^80 hhhhh20&&&&nd
Sample Output:100
*/
package main

import "strconv"

func parseStrToInt(s string) int64 {
	var buff string

	r := []rune(s)
	for _, r2 := range r {
		if !(r2 < '0' || r2 > '9') {
			buff += string(r2)
		}
	}
	res, _ := strconv.ParseInt(buff, 10, 64)
	return res
}

func adding(s1, s2 string) int64 {
	return parseStrToInt(s1) + parseStrToInt(s2)
}

func main() {
	adding("%^80", "hhhhh20&&&&nd")
}
