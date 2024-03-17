package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount = 1000
	var expectedReturn = 0.1
	var years = 10
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println(futureValue)

}
