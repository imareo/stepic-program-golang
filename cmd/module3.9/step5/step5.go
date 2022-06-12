/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и
дождаться результатов выполнения вызванных функций.
Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
*/
package main

import (
	"log"
	"sync"
	"time"
)

var d int

func main() {

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()

	if d != 10 {
		log.Fatalf("Error %d", d)
	}
	log.Print("All right")
}

func work() {
	time.Sleep(1 * time.Second)
	d++
}
