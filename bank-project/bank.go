package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalancefile = "balance.txt"

func main(){
	var accountBalance, err = getBalanceFromFile()

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("==========================")
		panic(err)
	}

	fmt.Println("Welcome to GO Bank!")

	for {
		appIntro()

		var userChoice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			fmt.Println("Your current account balance is: ", accountBalance)
		} else if userChoice == 2 {
			fmt.Print("Your Deposit Amount: ")
			var userDeposit float64
			fmt.Scan(&userDeposit)

			if userDeposit <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0.")
				continue
			}

			accountBalance = accountBalance + userDeposit
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance Updated! Latest Balance:", accountBalance)
		} else if userChoice == 3 {
			fmt.Print("Withdrawal Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount! Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw more than you have.")
				continue
			}

			accountBalance = accountBalance - withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance Updated! Latest Balance:", accountBalance)
		} else {
			fmt.Println("Goodbye!!")
			break
		}
	}
	fmt.Println("Thanks for choosing our Bank!")
}

func writeBalanceToFile(balance  float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalancefile, []byte(balanceText), 0644)
}
func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalancefile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}
func appIntro(){
	fmt.Println("================================")
	fmt.Println("How can I Help you?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
