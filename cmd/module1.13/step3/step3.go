/*
Дано трехзначное число. Переверните его, а затем выведите.
Sample Input:843
Sample Output:348
*/
package main

import "fmt"

func main() {
	var num string
	_, _ = fmt.Scan(&num)
	fmt.Println(reverse(num))
}

func reverse(num string) string {
	var res string
	for i := 0; i < len(num); i++ {
		res = string(num[i]) + res
	}
	return res
}
