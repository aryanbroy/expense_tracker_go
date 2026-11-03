package main

import (
	"fmt"
	"log"
	"strconv"
)

func fetchLatestId() int {
	data := ReadJsonFile(fileName)
	var latestId int

	for _, expense := range data {
		latestId = expense.Id
	}

	return latestId
}

func HandleExpenseAddition(args []string) {
	if len(args) <= 3 {
		log.Fatalln("Empty arguments passed")
	} else if len(args) > 4 {
		log.Fatalln("More than required arguments were passed")
	}
	desc := args[2]
	amountStr := args[3]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		log.Fatalln("Error converting string to int: ", err)
	}

	latestId := fetchLatestId()
	newId := latestId + 1

	newData := &Expense{
		Id:          newId,
		Description: desc,
		Amount:      amount,
	}

	AppendToJson(newData)
	fmt.Println("Successfully appended to json")
}
