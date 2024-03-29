package main

import "fmt"

func main() {
	prices := []float64{4.24, 5.55, 6.66}
	fmt.Println(prices[0:1])
	prices[1] = 5.99
	prices[2] = 39.99

	fmt.Println(prices)
	prices = append(prices, 9.99)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"tantuni"}
// 	productNames[2] = "lahmacun"

// 	prices := [4]float64{10.99, 5.99, 3.99, 7.99} //? 0 1 2 3
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[0])
// 	feauturedPrices := prices[1:] //? 	feauturedPrices := prices[1:3]
// 	feauturedPrices[0] = 1.99
// 	highlightedPrices := feauturedPrices[:2]
// 	fmt.Println(feauturedPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(feauturedPrices), cap(feauturedPrices))
// 	fmt.Println(len(prices), cap(prices))
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }
