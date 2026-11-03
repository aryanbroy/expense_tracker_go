package main

import (
	"fmt"
	"log"
	"strconv"
)

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
	newData := &Expense{
		Id:          1,
		Description: desc,
		Amount:      amount,
	}

	AppendToJson(newData)
	fmt.Println("Successfully appended to json")
}
