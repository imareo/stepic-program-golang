/*
Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на
стандартный ввод с новой строки (количество чисел не известно). Требуется прочитать все эти числа и вывести в
стандартный вывод их сумму.
Несколько подсказок: для чтения вы можете использовать типы bufio.Reader и bufio.Scanner, а для записи - bufio.Writer.
При чтении помните об ошибке io.EOF. Конвертирование данных из строки в целое число и обратно осуществляется функциями
Atoi и Itoa из пакета strconv соответственно. Чтение производится из стандартного ввода (os.Stdin), а запись - в
стандартный вывод (os.Stdout).
Все указанные в тексте задачи пакеты (strconv, bufio, os, io), уже импортированы (другие использовать нельзя), package
main объявлен.

Sample Input:
33
47
12
79
15
Sample Output:186
*/
package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// todo make test ?
func sumScanNumbers() (res int) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		chunk := scanner.Text()
		num, _ := strconv.Atoi(chunk)
		res += num
	}
	return
}

func main() {
	_, err := io.WriteString(os.Stdout, strconv.Itoa(sumScanNumbers()))
	if err != nil {
		return
	}
}
