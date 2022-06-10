package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("./task.data")
	defer file.Close()

	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		rs, _ := reader.ReadString(';')
		if string(rs) == "0;" {
			fmt.Printf("result: %d", i)
			break
		}
	}
}
