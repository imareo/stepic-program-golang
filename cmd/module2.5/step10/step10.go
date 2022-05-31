/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы
(считая с нуля)

Sample Input:ihgewlqlkot
Sample Output:hello
*/
package main

import "fmt"

func getEven(s string) string {
	var res string

	r := []rune(s)
	for i := 1; i < len(r); i += 2 {
		res += string(r[i])
	}
	return res
}

func main() {
	var str string

	_, _ = fmt.Scan(&str)
	fmt.Println(getEven(str))
}
