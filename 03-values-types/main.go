package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	// var investmentAmount float64
	// var expectedReturn float64
	// var years float64
	var investmentAmount, expectedReturn, years float64

	fmt.Println("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter the expected return: ")
	fmt.Scan(&expectedReturn)
	fmt.Println("Enter the years: ")
	fmt.Scan(&years)
	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
