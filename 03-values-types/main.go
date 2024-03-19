package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64 = 1000
	expectedReturn := 0.1
	years := 10.0
	futureValue := investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
