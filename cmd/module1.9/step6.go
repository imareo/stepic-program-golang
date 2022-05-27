/*
По данному трехзначному числу определите, все ли его цифры различны.
Формат входных данных
На вход подается одно натуральное трехзначное число.
Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:237
Sample Output 1:YES

Sample Input 2:117
Sample Output 2:NO
*/
package main

import "fmt"

func main() {
	var inp string

	_, _ = fmt.Scan(&inp)
	if inp[0] == inp[2] || inp[1] == inp[2] || inp[0] == inp[1] {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
