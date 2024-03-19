package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount, years float64 = 1000, 10
	expectedReturn := 0.1

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Println(futureValue)

}
