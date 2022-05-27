/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
Формат входных данных
На вход подается номер билета - одно шестизначное  число.
Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:613244
Sample Output:YES
*/
package main

import "fmt"

func main() {
	var inp string

	_, _ = fmt.Scan(&inp)
	sumF := int(inp[0]) + int(inp[1]) + int(inp[2])
	sumS := int(inp[3]) + int(inp[4]) + int(inp[5])
	if sumF == sumS {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
