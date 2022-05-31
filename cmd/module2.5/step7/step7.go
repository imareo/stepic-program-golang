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
