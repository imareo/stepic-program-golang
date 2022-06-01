/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

Sample Input:LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
Sample Output:L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	_, _ = fmt.Scan(&str)
	fmt.Println(strings.Trim(strings.ReplaceAll(str, "", "*"), "*"))
}
