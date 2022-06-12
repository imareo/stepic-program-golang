/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться
результатов ее выполнения.
Функция work() ничего не принимает и не возвращает.
*/
package main

import (
	"log"
	"time"
)

var d string

func main() {

	done := make(chan struct{})
	go func() {
		defer close(done)
		work()
	}()
	<-done

	if d != "nil" {
		log.Fatal("Error")
	}
	log.Print("All right")
}

func work() {
	time.Sleep(1 * time.Second)
	d = "nil"
}
