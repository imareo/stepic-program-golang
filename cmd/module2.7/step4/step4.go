/*Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

Sample Input:1112221112
Sample Output:2
*/
package main

import "fmt"

func maxDigit(num string) uint8 {
	var max uint8

	for i := range num {
		if num[i]-'0' > max {
			max = num[i] - '0'
		}
	}
	return max
}

func main() {
	var num string

	_, _ = fmt.Scan(&num)
	fmt.Println(maxDigit(num))
}
