package main

import (
	"log"
)

func fetchLatestId() int {
	data := ReadJsonFile(fileName)
	var latestId int

	for _, expense := range data {
		latestId = expense.Id
	}

	return latestId
}

func HandleExpenseCommands(args []string) {
	command := args[1]

	switch command {
	case "add":
		HandleExpenseAddition(args)
	default:
		log.Fatalln("Invalid command, use the ones specified in the README.md file")
	}
}
