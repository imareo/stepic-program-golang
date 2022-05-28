/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N,
в порядке возрастания.

Sample Input:50
Sample Output:1 2 4 8 16 32
*/
package main

import "fmt"

// todo add test ?
func get2Exp(num int) {
	for i := 1; i <= num; i *= 2 {
		fmt.Print(i, " ")
	}
}

func main() {
	var num int

	_, _ = fmt.Scan(&num)
	get2Exp(num)
}
