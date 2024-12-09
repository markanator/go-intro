package main

import (
	"example.com/markanator/banking-calc/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile, 1000)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("--------------------")
		//panic("Can't continue, sorry.")
	}
	for {
		fmt.Println("Welcome to your account!")
		fmt.Println("Contact us at: ", randomdata.PhoneNumber())
		PresentOptions()
		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
			return
		}

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
			fileops.WriteValueToFile(accountBalanceFile, accountBalance)
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
			fileops.WriteValueToFile(accountBalanceFile, accountBalance)
			fmt.Printf("Your new balance is %.2f\n", accountBalance)
		default:
			fmt.Println("Invalid choice")
			fmt.Println("Thanks for choosing Go Bank!")
			// ends app
			return
		}
	}
}
