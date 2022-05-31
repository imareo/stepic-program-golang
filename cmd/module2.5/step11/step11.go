/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести
получившуюся строку

Sample Input:zaabcbd
Sample Output:zcd
*/
package main

import (
	"fmt"
	"strings"
)

func delMultiChar(str string) string {
	var res []int32

	for _, i := range str {
		if strings.Count(str, string(i)) < 2 {
			res = append(res, i)
		}
	}
	return string(res)
}

func main() {
	var str string

	_, _ = fmt.Scan(&str)
	fmt.Println(delMultiChar(str))
}
