package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var tax_rate float64

	// fmt.Print("Enter your Renvenue: ")
	// fmt.Scan(&revenue)

	revenue := getUserInput("Enter your Revenue: ")
	expenses := getUserInput("Enter your Expenses: ")
	tax_rate := getUserInput("Enter the Tax Rate: ")

	// fmt.Print("Enter your Expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("Enter the Tax rate: ")
	// fmt.Scan(&tax_rate)

	ebt, profit, ratio := calculateFinances(revenue, expenses, tax_rate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f", ratio)

}

// Function Practice
func getUserInput(inputString string) float64 {
	var inputValue float64
	fmt.Print(inputString)
	fmt.Scan(&inputValue)
	return inputValue
}

func calculateFinances(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	var profit float64 = ebt * (1 - (tax_rate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio
}
