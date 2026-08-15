package main

import (
	"fmt"

	"example.com/bank/bankutils"
	"github.com/Pallinder/go-randomdata"
)

const accountBalancefile = "balance.txt"

func main(){
	var accountBalance, err = bankutils.GetFloatFromFile(accountBalancefile)

	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("==========================")
		panic(err)
	}

	fmt.Println("Welcome to GO Bank!")
	fmt.Println("Reach us 24/7: ",randomdata.PhoneNumber())

	// MAIN CLI LOOP
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
			bankutils.WriteBalanceToFile(accountBalance, accountBalancefile)
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
			bankutils.WriteBalanceToFile(accountBalance, accountBalancefile)
			fmt.Println("Balance Updated! Latest Balance:", accountBalance)
		} else {
			fmt.Println("Goodbye!!")
			break
		}
	}

	fmt.Println("Thanks for choosing our Bank!")
}
