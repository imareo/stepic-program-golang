/*Найдите самое большее число на отрезке от a до b, кратное 7.

Sample Input:100 500
Sample Output:497
*/
package main

import "fmt"

func solver(a int, b int) (int, bool) {
	for i := b; i > a; i-- {
		if i%7 == 0 {
			return i, true
		}
	}
	return 0, false
}

func main() {
	var a, b int

	_, _ = fmt.Scan(&a, &b)
	res, isAns := solver(a, b)
	if isAns {
		fmt.Println(res)
	} else {
		fmt.Println("NO")
	}
}
