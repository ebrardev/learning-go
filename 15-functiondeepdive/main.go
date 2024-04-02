package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := doubleNumbers(&numbers)

	fmt.Println(doubled)

}

func doubleNumbers(numbers *[]int) []int {
	dbNumbers := []int{}
	for _, val := range *numbers {

		dbNumbers = append(dbNumbers, val*2)

	}
	return dbNumbers
}
