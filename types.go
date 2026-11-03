package main

import "time"

type Expense struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}
