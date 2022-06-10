package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Students struct {
	ID       int    `json:"ID"`
	Number   string `json:"Number"`
	Year     int    `json:"Year"`
	Students []struct {
		LastName   string `json:"LastName"`
		FirstName  string `json:"FirstName"`
		MiddleName string `json:"MiddleName"`
		Birthday   string `json:"Birthday"`
		Address    string `json:"Address"`
		Phone      string `json:"Phone"`
		Rating     []int  `json:"Rating"`
	} `json:"Students"`
}

type Result struct {
	Average float32 `json:"Average"`
}

func main() {

	var students Students
	var result Result

	var resNumbers int

	// data := readFile()  for test
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &students)
	if err != nil {
		return
	}
	for _, student := range students.Students {
		resNumbers += len(student.Rating)
	}

	result.Average = float32(resNumbers) / float32(len(students.Students))
	res, _ := json.MarshalIndent(&result, "", "    ")

	fmt.Printf("%s", res)
}

func readFile() []byte {
	file, err := os.Open("./test.json")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)

	return data
}
