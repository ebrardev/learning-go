package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User
	appUser = User{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}

	outputUserDetails(appUser)
}

func outputUserDetails(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createAt)

}
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
