package main

import (
	"fmt"
	"os"
)

func main() {
	var accountBalance float64 = 1500.00
	fmt.Println("Welcome to GoWriteBank!!!!!!")

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Funds")
		fmt.Println("3. Withdraw Funds")
		fmt.Println("4. Exit")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("Your choice: ")
		switch choice {
		case 1:
			fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Println("Enter amount to deposit: ")
			fmt.Scanln(&depositAmount)
			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Printf("Successfully deposited $%.2f. New balance: $%.2f\n", depositAmount, accountBalance)
			} else {
				fmt.Println("Invalid deposit amount.")
			}
		case 3:
			var withdrawAmount float64
			fmt.Println("Enter amount to withdraw: ")
			fmt.Scanln(&withdrawAmount)
			if withdrawAmount > 0 && withdrawAmount <= accountBalance {
				accountBalance -= withdrawAmount
				fmt.Printf("Successfully withdrew $%.2f. New balance: $%.2f\n", withdrawAmount, accountBalance)
			} else {
				fmt.Println("Invalid withdrawal amount or insufficient funds.")
			}
		default:
			// Write the final balance to a file before exiting
			file, err := os.Create("final_balance.txt")
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer file.Close()

			_, err = file.WriteString(fmt.Sprintf("Final account balance: $%.2f\n", accountBalance))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
			fmt.Println("Thank you for banking with GoWriteBank. Goodbye!")
			return
		}
	}
}
