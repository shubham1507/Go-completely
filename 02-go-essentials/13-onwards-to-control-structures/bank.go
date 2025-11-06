package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.00

	//for i := 0; i < 50; i++
	for {
		fmt.Println("Welcome to the Bank of Go!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Printf("Your balance is $%.2f\n", accountBalance)
		} else if choice == 2 {
			var amount float64
			fmt.Println("Enter amount to deposit:")
			fmt.Scan(&amount)
			accountBalance += amount
			fmt.Printf("You have deposited $%.2f and now your account balance is $%.2f\n", amount, accountBalance)
		} else if choice == 3 {
			var amount float64
			fmt.Println("Enter amount to withdraw:")
			fmt.Scan(&amount)
			if amount > accountBalance {
				fmt.Println("Insufficient funds!")
				return
			}
			accountBalance -= amount
			fmt.Printf("You have withdrawn $%.2f\n and now the balance is $%.1f", amount, accountBalance)
		} else if choice == 4 {
			fmt.Println("Thank you for banking with us!")
			break
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}

	}

}
