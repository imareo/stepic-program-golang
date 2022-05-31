/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная
строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная -
вывести Right иначе - вывести Wrong

Sample Input:Быть или не быть.
Sample Output:Right
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(solver(text))
}

func solver(text string) string {
	textAsRune := []rune(text)
	if unicode.IsUpper(textAsRune[0]) && textAsRune[len(textAsRune)-1] == '.' {
		return "Right"
	}
	return "Wrong"
}
