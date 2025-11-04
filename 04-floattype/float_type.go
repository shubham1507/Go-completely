package main

import "fmt"

func main() {
	luckyNumber := 11
	evenMoreLuckyNumber := luckyNumber * 2

	var myNumber float64 = float64(evenMoreLuckyNumber) / 4
	var smallerFloat float32 = 1.234567

	fmt.Println(myNumber)
	fmt.Println(smallerFloat)
}
