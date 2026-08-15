package bankutils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Add Balance to File
func WriteBalanceToFile(balance  float64, fileName string){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}


// Get Balance from File
func GetFloatFromFile(fileName string) (float64, error){
	data, err := os.ReadFile(fileName)


	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	valueText := string(data)
	if valueText == "" {
		return 1000, nil // no balance stored yet, default to 1000
	}

	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return value, nil
}
