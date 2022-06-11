package main

import (
	"fmt"
)

func task2(c chan string, s string) {
	c <- s + " "
	c <- s + " "
	c <- s + " "
	c <- s + " "
	c <- s + " "
}

func main() {
	ch := make(chan string, 5)

	go task2(ch, "h2o")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	close(ch)
}
