package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	// fmt.Print("Enter your Renvenue: ")
	// fmt.Scan(&revenue)

	fmt.Print("Enter your Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the Tax rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	var profit float64 = ebt * (1 - (tax_rate / 100))

	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

// Function Practice
func getUserInput(inputString string) float64 {
	var inputValue float64

	fmt.Print(inputString)
	fmt.Scan(&inputValue)

	return inputValue
}
