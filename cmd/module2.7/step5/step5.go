/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести
получившееся число.

Sample Input:9119
Sample Output:811181
*/
package main

import (
	"fmt"
	"strconv"
)

func squaredDigit(num string) string {
	var res string

	for _, i := range num {
		sqDig := strconv.Itoa(int(i-'0') * int(i-'0'))
		res += sqDig
	}
	return res
}

func main() {
	var num string

	_, _ = fmt.Scan(&num)
	fmt.Println(squaredDigit(num))
}
