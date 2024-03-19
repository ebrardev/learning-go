package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxtRate float64

	outputText("Enter the revenue: ")
	fmt.Scan(&revenue)
	outputText("Enter the expenses: ")
	fmt.Scan(&expenses)
	outputText("Enter the tax rate: ")
	fmt.Scan(&taxtRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxtRate/100)
	ratio := profit / revenue
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}

func outputText(text string) {
	fmt.Println(text)
}
