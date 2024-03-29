package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"tantuni"}
	productNames[2] = "lahmacun"

	prices := [4]float64{10.99, 5.99, 3.99, 7.99} //? 0 1 2 3
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[0])
	feauturedPrices := prices[1:] //? 	feauturedPrices := prices[1:3]
	feauturedPrices[0] = 1.99
	highlightedPrices := feauturedPrices[:2]
	fmt.Println(feauturedPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(feauturedPrices), cap(feauturedPrices))
	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

}
