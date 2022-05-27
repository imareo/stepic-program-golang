/*
Требуется написать программу, при выполнении которой с клавиатуры считываются
два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел
от A до B  включительно.

Sample Input:1 5
Sample Output:15
*/
package main

import "fmt"

func main() {
	var inp1, inp2, res int
	_, _ = fmt.Scan(&inp1, &inp2)
	for i := inp1; i < inp2+1; i++ {
		res += i
	}
	fmt.Println(res)
}
