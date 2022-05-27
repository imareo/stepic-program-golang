/*
Напишите программу, которая в последовательности чисел находит сумму двузначных чисел,
кратных 8. Программа в первой строке получает на вход число n - количество чисел в
последовательности, во второй строке -- n чисел, входящих в данную последовательность.

Sample Input:5
38 24 800 8 16
Sample Output:40
*/
package main

import "fmt"

func main() {
	var inp1, inp2, res int
	_, _ = fmt.Scan(&inp1)
	for i := 1; i < inp1+1; i++ {
		_, _ = fmt.Scan(&inp2)
		if inp2%8 == 0 && inp2 >= 10 && inp2 < 100 {
			res += inp2
		}
	}
	fmt.Println(res)
}
