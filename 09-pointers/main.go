package main

import "fmt"

func main() {

	age := 30

	fmt.Println("Age is: ", age)
	adultYears := getAdultYears(age)
	fmt.Println("Adult years: ", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
