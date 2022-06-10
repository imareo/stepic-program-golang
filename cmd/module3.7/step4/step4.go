/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же
формате.
Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести
на стандартный вывод в том же формате.

Sample Input:2020-05-15 08:00:00
Sample Output:2020-05-15 08:00:00
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func solveTime(t string) string {
	const timeLayout = "2006-01-02 15:04:05"

	it, err := time.Parse(timeLayout, t)
	if err != nil {
		panic(err)
	}
	if it.Hour() < 13 {
		return it.Format(timeLayout)
	} else {
		return it.AddDate(0, 0, 1).Format(timeLayout)
	}
}

func main() {

	dt := bufio.NewScanner(os.Stdin)
	dt.Scan()
	fmt.Print(solveTime(dt.Text()))
}
