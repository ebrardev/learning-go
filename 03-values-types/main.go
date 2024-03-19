package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64 = 1000
	expectedReturn := 0.1
	years := 10.0
	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	fmt.Println(futureValue)

}
