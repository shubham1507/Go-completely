package main

import "fmt"

// function returns two values: sum and product
func calculate(a int, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	product := a * b
	return sum, diff, product
}

func main() {
	s, d, p := calculate(6, 4)
	fmt.Println("Sum =", s)
	fmt.Println("Difference =", d)
	fmt.Println("Product =", p)
}
