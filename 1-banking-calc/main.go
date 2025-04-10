package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Could not read file")
	}
	// convert byte[] into string
	balanceString := string(data)
	// convert string to float
	balanceFloat, err := strconv.ParseFloat(balanceString, 64)
	if err != nil {
		return 1000, errors.New("Could not parse balance")
	}
	// only return float
	return balanceFloat, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	for {
		accountBalance, err := getBalanceFromFile()
		if err != nil {
			fmt.Println("ERROR!")
			fmt.Println(err)
			fmt.Println("--------------------")
			//panic("Can't continue, sorry.")
		}
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is %.2f\n", accountBalance)
		case 2:
			fmt.Println("Amount to Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Error: Invalid amount.")
				continue
			}
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		case 3:
			fmt.Println("Amount to Withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("Error: Invalid amount.")
				continue
			}
			if withdrawalAmount >= accountBalance {
				fmt.Println("Error: Insufficient funds.")
				continue
			}
			accountBalance -= withdrawalAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thanks for choosing Go Bank!")
			// ends app
			return
		}
	}
}
