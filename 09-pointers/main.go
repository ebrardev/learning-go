package main

import "fmt"

func main() {

	age := 30
	var agePointer *int
	agePointer = &age

	fmt.Println("Age is: ", *agePointer)
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
