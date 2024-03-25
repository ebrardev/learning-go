package main

import "fmt"

func main() {

	age := 30
	var agePointer *int
	agePointer = &age

	fmt.Println("Age is: ", *agePointer)
	// adultYears := getAdultYears(age)
	// fmt.Println("Adult years: ", adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}
