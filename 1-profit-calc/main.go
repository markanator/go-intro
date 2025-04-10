package main

import (
	"errors"
	"fmt"
	"os"
)

const accountBalanceFile = "reports.txt"

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile(accountBalanceFile, []byte(results), 0644)
}

func main() {
	revenue, err := getUserInput("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	// ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Ebt: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	input, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, errors.New("unable to read user input")
	}
	if input <= 0 {
		return 0, errors.New("must be positive number")
	}
	return float64(input), nil
}
