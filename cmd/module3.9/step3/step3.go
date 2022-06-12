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
