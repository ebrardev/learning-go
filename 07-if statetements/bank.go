package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to the Bank of Golang")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)



	wantsCheckBalance:= choice == 1
     
	  if wantsCheckBalance {

	  }
	fmt.Println("Your choice is: ", choice)
 
	
}
