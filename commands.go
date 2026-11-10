package main

import (
	"log"
)

func HandleExpenseCommands(args []string) {
	command := args[1]

	switch command {
	case "add":
		// latestId := FetchLatestId(fileName)
		// fmt.Println("Latest id: ", latestId)
		HandleExpenseAddition(args)
	case "list":
		HandleExpensesList()
	case "summary":
		HandleSummary()
	case "delete":
		HandleExpenseDel(args)
	default:
		log.Fatalln("Invalid command, use the ones specified in the README.md file")
	}
}
