package main

import (
	"fmt"

	"example.com/ebrarmea/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error reading balance from file: ", err)
		fmt.Println("------------------------------------")
		fmt.Println(err)
		fmt.Println("------------------------------------")
		// panic("cant continue sorry")

	}
	var depositAmount, WithdrawAmount float64
	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("Hello, ", randomdata.FullName(randomdata.RandomGender))
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		} else if choice == 3 {
			fmt.Println("Withdraw")
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&WithdrawAmount)
			if WithdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
			} else {
				accountBalance -= WithdrawAmount // accountBalance = accountBalance - WithdrawAmount
				fmt.Println("Balance updated!, Your account balance is: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
