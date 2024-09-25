package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	var years float64
	expectedReturn := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Returns: ")
	fmt.Scan(&expectedReturn)

	futureValue := (investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
