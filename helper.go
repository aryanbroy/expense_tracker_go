package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
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

	err = checkFile(fileName)
	if err != nil {
		log.Fatalln("Error validating file:  ", err)
	}

	latestId := fetchLatestId()
	newId := latestId + 1

	newData := &Expense{
		Id:          newId,
		Description: desc,
		Amount:      amount,
		Date:        time.Now(),
	}

	AppendToJson(newData)
	fmt.Println("Successfully appended to json")
}

func HandleExpensesList() {
	err := checkFile(fileName)
	if err != nil {
		log.Fatalln("Error validating file: ", err)
	}

	expenses := ReadJsonFile(fileName)

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]any{"ID", "Date", "Description", "Amount"})

	for _, expense := range expenses {
		t := expense.Date
		formattedTime := t.Format("2006-01-02")
		table.Append([]any{expense.Id, formattedTime, expense.Description, expense.Amount})
	}

	table.Render()
}
