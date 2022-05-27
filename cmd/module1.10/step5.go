/*
Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0
(само число 0 в последовательность не входит, а служит как признак ее окончания).

Sample Input:1 3 3 1 0
Sample Output:2
*/
package main

import "fmt"

func main() {
	var inp, max, res int
	for _, _ = fmt.Scan(&inp); inp != 0; _, _ = fmt.Scan(&inp) {
		if inp > max {
			res = 1
			max = inp
		} else if inp == max {
			res += 1
		}
	}
	fmt.Println(res)
}
