/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Sample Input:4 5 6 7
Sample Output:4
*/
package main

import "fmt"

// todo test ?
func minimumFromFour() int {
	var num, min int
	isFirst := true

	for i := 0; i < 4; i++ {
		switch _, _ = fmt.Scan(&num); {
		case isFirst:
			min, isFirst = num, false
		case num < min:
			min = num
		}
	}
	return min
}
