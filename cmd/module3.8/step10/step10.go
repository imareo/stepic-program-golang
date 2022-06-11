/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап
конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы
должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд.
Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()
*/
package main

import (
	"fmt"
	"time"
)

func removeDuplicates(i, o chan string) {
	var acc string

	for x := range i {
		if x != acc {
			acc = x
			o <- x
		}
	}
	close(o)
}

func main() {
	in := make(chan string)
	out := make(chan string)
	go removeDuplicates(in, out)
	go func() {
		defer close(in)
		for _, r := range "112334456" {
			in <- string(r)
		}
		time.Sleep(2 * time.Second)
		for _, r := range "77889900" {
			in <- string(r)
		}
	}()

	go func() {
		defer close(in)
		time.Sleep(1 * time.Second)
		for _, r := range "aabbsdd" {
			in <- string(r)
		}
		time.Sleep(3 * time.Second)
		for _, r := range "mmnnopp" {
			in <- string(r)
		}
	}()

	for x := range out {
		fmt.Print(x)
	}
}
