/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
Например, если k=13257=3*3600+40*60+57, то h=3 и m=40.
На вход программе подается целое число k (0 < k < 86399).
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/
package main

import "fmt"

func getHourMinute(k int) (h, m int) {
	h = k / 3600
	m = (k - h*3600) / 60
	return h, m
}

func main() {
	var k int

	_, _ = fmt.Scan(&k)
	h, m := getHourMinute(k)
	fmt.Printf("It is %d hours %d minutes.", h, m)
}
