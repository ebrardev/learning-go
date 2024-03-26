package main

import (
	"fmt"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	outputUserDetails(lastName, firstName, birthdate)
}

func outputUserDetails(firstName, lastName, birthdate string) {
	fmt.Println("User Details:")
	fmt.Println("First Name: ", firstName)
	fmt.Println("Last Name: ", lastName)
	fmt.Println("Birthdate: ", birthdate)
}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
