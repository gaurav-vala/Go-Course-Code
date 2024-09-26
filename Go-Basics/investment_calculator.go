package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64 = 1000
	var years float64
	expectedReturn := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Investment Years: ")
	outputText("Investment Years:")
	fmt.Scan(&years)

	fmt.Print("Expected Returns: ")
	fmt.Scan(&expectedReturn)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturn, years)

	// futureValue := (investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value : %.0f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (Adjusted for Inflation): %0.f\n", futureRealValue)

	// fmt.Printf(`Future Value : %v
	// Future Value (Adjusted for Inflation): %v`, futureValue, futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println(futureRealValue)

	fmt.Print(formattedFV, formattedFRV)

}

func outputText(text1 string) {
	fmt.Print(text1)
}

func calculateFutureValue(investmentAmount, expectedReturn, years float64) (float64, float64) {
	fv := (investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
