package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalanceFiles(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func readFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parsed stored balance value")
	}

	return balance, nil
}

func main() {
	var accountBalance, err = readFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("===============")
	}

	fmt.Println("Welcome to GO Bank!")
	for {
		fmt.Println("===============")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Println("===============")

		var choice int
		fmt.Print("Your Choice ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		// switch choice {
		// case 1:
		// 	fmt.Println("Your Balance is", accountBalance)
		// case 2:
		// 	fmt.Print("Your Deposit: ")
		// case 3:
		// 	fmt.Print("Withdrawal Amount:")
		// default:
		// 	fmt.Print("Goodbye!")
		// }

		if choice == 1 {
			fmt.Println("Your Balance is", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("The Deposit Amount is Invalid!")
				continue
			}
			accountBalance += depositAmount // same as accountBalance = accountBalance + depositAmount
			fmt.Println("Balance Updated! New Balance:", accountBalance)
			writeBalanceFiles(accountBalance)
		} else if choice == 3 {
			fmt.Print("Withdrawal Amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount <= 0 {
				fmt.Println("The Deposit Amount is Invalid!")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid Amount! You can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Money Withdrawn! New Balance:", accountBalance)
			writeBalanceFiles(accountBalance)
		} else {
			fmt.Print("Goodbye!")
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank!")

}
