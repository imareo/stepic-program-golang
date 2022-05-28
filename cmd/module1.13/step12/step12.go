/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений:
"n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova
или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять
ровно один пробел.
*/
package main

import "fmt"

func getNumeral(num int) string {
	var res string

	switch {
	case num < 20 && num > 10:
		res = fmt.Sprintf("%d korov", num)
	case num%10 == 1:
		res = fmt.Sprintf("%d korova", num)
	case num%10 > 1 && num%10 < 5:
		res = fmt.Sprintf("%d korovy", num)
	default:
		res = fmt.Sprintf("%d korov", num)
	}
	return res
}

func main() {
	var num int

	_, _ = fmt.Scan(&num)
	fmt.Println(getNumeral(num))
}
