package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%f", balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance float64 = 1000.0
	var depositAmount, WithdrawAmount float64
	fmt.Println("Welcome to the Bank of Golang")
	for i := 0; i < 3; i++ {

		fmt.Println("What would you like to do today?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance:= choice == 1

		switch choice {
		case 1:
			fmt.Println("Checking your balance")
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Println("Deposit")
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

		}

		if choice == 1 {
			fmt.Println("Checking your balance")
			fmt.Println("Your account balance is: ", accountBalance)

		} else if choice == 2 {
			fmt.Println("Deposit")
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount")

				continue
			}
			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println(" Balance updated!, Your account balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)

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
			break

		}
	}

	fmt.Println("Thank you for using the Bank of Golang")

}
