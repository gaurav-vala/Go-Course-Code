package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate User Input
//  => Show error & exit if invalid input is provided
// 		- No Negative numbers
// 		- Not 0
//  2) Store Calculated results into file

func main() {
	// var revenue float64
	// var expenses float64
	// var tax_rate float64

	// fmt.Print("Enter your Renvenue: ")
	// fmt.Scan(&revenue)

	for {
		revenue, err1 := getUserInput("Enter your Revenue: ")
		// os.WriteFile("profit.txt", []byte(revenue), 0644)

		if err1 != nil {
			fmt.Println(err1)
			fmt.Println("==========")
			return
		}

		expenses, err2 := getUserInput("Enter your Expenses: ")

		if err2 != nil {
			fmt.Println(err2)
			fmt.Println("==========")
			return
		}

		tax_rate, err3 := getUserInput("Enter the Tax Rate: ")

		if err3 != nil {
			fmt.Println(err3)
			fmt.Println("==========")
			return
		}

		// if err1 != nil || err2 != nil || err3 != nil {
		// 	fmt.Println(err1)
		// 	fmt.Println("==========")
		// 	return
		// }

		// fmt.Print("Enter your Expenses: ")
		// fmt.Scan(&expenses)
		// fmt.Print("Enter the Tax rate: ")
		// fmt.Scan(&tax_rate)

		ebt, profit, ratio := calculateFinances(revenue, expenses, tax_rate)

		fmt.Printf("%.1f\n", ebt)
		fmt.Printf("%.1f\n", profit)
		fmt.Printf("%.1f", ratio)

		storeProfit(ebt, profit, ratio)
	}

}

func storeProfit(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}

// Function Practice
func getUserInput(inputString string) (float64, error) {
	var inputValue float64
	fmt.Print(inputString)
	fmt.Scan(&inputValue)
	if inputValue <= 0 {
		return 0, errors.New(">>> Input value is invalid")
	}
	return inputValue, nil
}

func calculateFinances(revenue, expenses, tax_rate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	var profit float64 = ebt * (1 - (tax_rate / 100))
	ratio := ebt / profit
	return ebt, profit, ratio
}
