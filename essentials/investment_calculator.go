package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64 = 10

	// Get Investment Amount from User - Scane the values from the Terminal
	outputText("Enter Your Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	// Calculate the values
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// Print the Value
	// fmt.Println(futureValue)
	fmt.Printf("Future Value: %.2f \nFuture Value (adjusted got Inflation): %.2f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted got Inflation): ",futureRealValue)
}

func outputText(text string, ) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
