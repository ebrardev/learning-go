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

func (u User) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createAt)

}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""

}

func newUser(firstName string, lastName string, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *User
	appUser = newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
