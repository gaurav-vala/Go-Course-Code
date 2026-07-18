package main

import (
	"errors"
	"fmt"
	"os"
)

// GOALS
// 1. Validate the user input
// 	=> Show error message & exit if invalid input is provided
// 	- No negative numbers
// 	- Not 0
// 2. Store Calculated results into file

func main (){

	revenue, err := getValueFromUser("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getValueFromUser("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getValueFromUser("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}


	// Calculate Earnings Before Tax (EBT) and Eearnings After Tax (PROFIT), Ratio Between EBT and PROFIT
	ebt, profit, ratio := calculateValues(revenue, expenses, taxRate)

	// Output EBT, profit and the ratio
	fmt.Println("Your Earning before Tax: ", ebt)
	fmt.Println("Your Profit: ", profit)
	fmt.Println("The Ratio: ", ratio)

	storeResults(ebt, profit, ratio)
}

func getValueFromUser(text string) (float64,error)  {
	var userVal float64

	fmt.Print(text)
	fmt.Scan(&userVal)

	if userVal <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}
	return userVal, nil
}

func calculateValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	// Calculate Earnings Before Tax (EBT) and Eearnings After Tax (PROFIT), Ratio Between EBT and PROFIT
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("profit-results.txt", []byte(results), 0644)
}
