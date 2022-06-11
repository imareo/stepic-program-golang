package main

import (
	"fmt"
)

func task(c chan int, n int) {
	c <- n + 1
}

func main() {
	ch := make(chan int, 1)

	go task(ch, 6)
	fmt.Println(<-ch)

	close(ch)
}
