/*
Из натурального числа удалить заданную цифру.

Sample Input:38012732 3
Sample Output:
801272
*/
package main

import (
	"fmt"
	"strings"
)

func delNum(num, del string) string {
	return strings.ReplaceAll(num, del, "")
}

func main() {
	var num, del string

	_, _ = fmt.Scan(&num, &del)
	fmt.Println(delNum(num, del))
}
