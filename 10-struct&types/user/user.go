package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func (u User) outputUserDetails() {
	fmt.Println(u.FirstName, u.lastName, u.birthdate, u.createAt)

}

func (u *User) clearUserName() {
	u.FirstName = ""
	u.lastName = ""

}

func newUser(firstName string, lastName string, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New(" ERROR ! All fields are required")
	}

	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}, nil
}
