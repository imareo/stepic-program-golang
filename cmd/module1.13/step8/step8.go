/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.
Sample Input:5
1 8 100 0 12
Sample Output:1
*/
package main

import "fmt"

// todo add test ?
func main() {
	var num, temp, res int

	_, _ = fmt.Scan(&num)
	for i := 0; i < num; i++ {
		_, _ = fmt.Scan(&temp)
		if temp == 0 {
			res++
		}
	}
	fmt.Println(res)
}
