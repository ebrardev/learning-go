package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dbNumbers := []int{}
	for _, val := range *numbers {

		dbNumbers = append(dbNumbers, transform(val))

	}
	return dbNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
