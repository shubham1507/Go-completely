/*
Problem Statement

Create a Go program that calculates financial metrics for a business based on user input.

The program should:

Ask the user to enter three values:

Revenue (total money earned)

Expenses (total cost spent)

Tax Rate (in percentage)

Using these values, calculate:

EBT (Earnings Before Tax) = revenue − expenses

Profit (after tax) = EBT × (1 − taxRate/100)

Profit Ratio = EBT / profit

Display the calculated results with proper formatting:

EBT with 1 decimal

Profit with 1 decimal

Ratio with 3 decimals

Use functions in a clean, modular way:

getUserInput() should read input from user

calculateFinancials() should calculate and return multiple values (EBT, Profit, Ratio)

Expected Output Example
Revenue: 10000
Expenses: 4500
Tax Rate: 18
5500.0
4510.0
1.220

Learning Goals

Practice taking user input in Go

Apply arithmetic / percentage formulas

Use a function that returns multiple values

Format output using Printf
*/

package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue")
	expenses := getUserInput("Expenses")
	taxRate := getUserInput("Tax Rate")
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(prompt string) float64 {
	var value float64
	fmt.Printf("%s: ", prompt)
	fmt.Scan(&value)
	return value
}
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
