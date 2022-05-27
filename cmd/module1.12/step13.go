/*
Напишите программу, принимающая на вход число N(N≥4), а затем NN целых чисел
- элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:5
41 -231 24 49 6
Sample Output:49
*/
package main

import "fmt"

func main() {
	var num int

	_, _ = fmt.Scan(&num)
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		_, _ = fmt.Scan(&slice[i])
	}
	fmt.Printf("%d ", slice[3])
}
