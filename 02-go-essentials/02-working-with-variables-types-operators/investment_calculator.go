package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount int
	const inflationRate = 2.0
	var expectedReturnRate float64
	var year int

	fmt.Print("investmentAmount:  ")
	fmt.Scan(&investmentAmount)
	fmt.Print("expectedReturnRate:  ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("year:  ")
	fmt.Scan(&year)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(year))
	futureValueAdjustedForInflation := futureValue / math.Pow(1+inflationRate/100, float64(year))

	formattedFV := fmt.Sprintf("Future value: %.1f \n", futureValue)
	formattedRFV := fmt.Sprintf("Future value adjusted for inflation: %.1f \n", futureValueAdjustedForInflation)

	fmt.Println(futureValueAdjustedForInflation)
	fmt.Println(futureValue)
	fmt.Print(formattedFV, formattedRFV)
}
