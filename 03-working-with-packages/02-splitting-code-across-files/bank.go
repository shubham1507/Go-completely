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
		return 1000, errors.New("Failed to find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse balance.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprintf("%.2f", balance)
	err := os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	if err != nil {
		return errors.New("Failed to write balance to file.")
	}
	return nil
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Welcome to Go Bank")
	fmt.Printf("Current balance: $%.2f\n", accountBalance)

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Printf("Deposited: $%.2f\n", depositAmount)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient funds.")
			} else {
				accountBalance -= withdrawAmount
				fmt.Printf("Withdrew: $%.2f\n", withdrawAmount)
			}
		case 4:
			err := writeBalanceToFile(accountBalance)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Balance saved. Exiting.")
			}
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
