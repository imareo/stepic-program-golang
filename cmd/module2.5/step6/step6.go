package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var en = "english"
	var ru = "русский"
	var enru = "english русский"

	// utf8.RuneCountString()
	fmt.Println(len(en), len(ru), len(enru))
	fmt.Println(utf8.RuneCountInString(en), utf8.RuneCountInString(ru), utf8.RuneCountInString(enru))
}
