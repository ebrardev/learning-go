package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.0
	var depositAmount float64 = 0.0
	var WithdrawAmount float64
	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance:= choice == 1

	if choice == 1 {
		fmt.Println("Checking your balance")
		fmt.Println("Your account balance is: ", accountBalance)

	} else if choice == 2 {
		fmt.Println("Deposit")
		fmt.Print("Enter the amount you want to deposit: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
		fmt.Println(" Balance updated!, Your account balance is: ", accountBalance)

	} else if choice == 3 {
		fmt.Println("Withdraw")
		fmt.Print("Enter the amount you want to withdraw: ")
		fmt.Scan(&WithdrawAmount)
		if WithdrawAmount > accountBalance {
			fmt.Println("Insufficient balance")
		} else {
			accountBalance -= WithdrawAmount // accountBalance = accountBalance - WithdrawAmount
			fmt.Println("Balance updated!, Your account balance is: ", accountBalance)
		}

	} else if choice == 4 {
		fmt.Println("Exiting the Bank of Golang")
	} else {
		fmt.Println("Invalid choice")

	}
}
