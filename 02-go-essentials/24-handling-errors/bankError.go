package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accounBalanceFile = "account_balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accounBalanceFile)
	if err != nil {
		return 1000, errors.New("could not read balance file, initializing with default balance of 1000")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0, errors.New("invalid balance format in file")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accounBalanceFile, []byte(balanceText), 0644)
}

func main() {
	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Welcome! Your current balance is: %.2f\n", balance)

	for {
		var choice int
		fmt.Println("Choose an option: 1) Deposit 2) Withdraw 3) Exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Deposit amount must be positive.")
				continue
			}
			balance += amount
			fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, balance)
		case 2:
			var amount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Withdrawal amount must be positive.")
				continue
			}
			if amount > balance {
				fmt.Println("Insufficient funds for this withdrawal.")
				continue
			}
			balance -= amount
			fmt.Printf("Withdrew %.2f. New balance: %.2f\n", amount, balance)
		case 3:
			writeBalanceToFile(balance)
			fmt.Println("Exiting. Balance saved.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
