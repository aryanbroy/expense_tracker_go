package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var fileName = "expense.json"

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	flag := false

	if err != nil {
		if os.IsNotExist(err) {
			flag = true
			_, err := os.Create(filename)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	fmt.Println("File exists: ", flag)
	return nil
}

func ReadJsonFile(fileName string) []Expense {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}

	data := []Expense{}

	json.Unmarshal(file, &data)

	return data
}

func AppendToJson(newData *Expense) {
	data := ReadJsonFile(fileName)

	data = append(data, *newData)

	dataBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalln("Error marshaling data: ", err)
	}

	err = os.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Fatalln("Error writing to json: ", err)
	}
}
