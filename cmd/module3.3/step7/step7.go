package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(forTest(100))
}

func forTest(num uint) uint {
	// after
	fn := func(num uint) uint {
		res := ""
		numAsString := strconv.FormatUint(uint64(num), 10)
		for pos := range numAsString {
			if (numAsString[pos]-'0') == 0 || (numAsString[pos]-'0')%2 != 0 {
				continue
			}
			res += string(numAsString[pos])
		}
		resOut, _ := strconv.ParseUint(res, 10, 64)
		if resOut == 0 {
			return 100
		}
		return uint(resOut)
	}
	// before
	return fn(num)
}
