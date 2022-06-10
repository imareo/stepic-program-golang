/*
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные, которые
вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории
https://github.com/semyon-dev/stepik-go/tree/master/work_with_json

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать
сумму полей global_id всех элементов, закодированных в файле.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type workSchema []struct {
	GlobalId int `json:"global_id"`
}

func main() {
	url := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var dataset workSchema
	var res int

	err = json.NewDecoder(resp.Body).Decode(&dataset)
	for _, item := range dataset {
		res += item.GlobalId
	}
	fmt.Println(res)
}
