package main

import "fmt"

func main() {
	var accountBalance float64 = 1500.00
	fmt.Println("Welcome to SwitchBank!!!!!!")

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
			fmt.Println("Thank you for banking with SwitchBank. Goodbye!")
			return
		}
	}
}
