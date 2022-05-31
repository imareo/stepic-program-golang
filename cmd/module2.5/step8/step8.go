/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является
палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях
(например, "топот", "око", "заказ").

Sample Input:топот
Sample Output:Палиндром
*/
package main

import "fmt"
import "bufio"
import "os"
import "strings"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(text string) bool {
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)
	return text == Reverse(text)
}

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if isPalindrome(text) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
