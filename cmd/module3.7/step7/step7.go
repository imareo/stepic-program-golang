/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей
датами.

Sample Input:13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:24h0m0s
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func td(t1 time.Time, t2 time.Time) string {
	dif := t1.Sub(t2)
	if dif < 0 {
		dif = -dif
	}
	return dif.String()
}

func timeDiff(t string) string {
	const timeLayout = "02.01.2006 15:04:05"
	var ta []string

	ta = strings.Split(t, ",")
	t1, _ := time.Parse(timeLayout, ta[0])
	t2, _ := time.Parse(timeLayout, ta[1])

	return td(t1, t2)
}

func main() {
	dt := bufio.NewScanner(os.Stdin)
	dt.Scan()
	fmt.Println(timeDiff(dt.Text()))
}
