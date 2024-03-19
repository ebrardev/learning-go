package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxtRate float64

	fmt.Println("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Println("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Println("Enter the tax rate: ")
	fmt.Scan(&taxtRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxtRate/100)
	ratio := profit / revenue
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
