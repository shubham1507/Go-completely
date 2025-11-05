package main

import "fmt"

// SI = (P * R * T) / 100

func simpleInterest(principal float64, rate float64, time float64) float64 {
	return (principal * rate * time) / 100
}

func main() {
	var principalAmount float64
	fmt.Println("Simple Interest Calculation")
	fmt.Print(" Enter Principal Amount: \n")
	fmt.Scan(&principalAmount)
	var rateOfInterest float64
	fmt.Print(" Enter Rate of Interest: \n")
	fmt.Scan(&rateOfInterest)
	var timePeriod float64
	fmt.Print(" Enter Time Period (in years): \n")
	fmt.Scan(&timePeriod)

	si := simpleInterest(principalAmount, rateOfInterest, timePeriod)
	fmt.Printf(" Simple Interest is: %.2f\n", si)
}
