package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := [4]float64{10.99, 5.99, 3.99, 7.99}
	fmt.Println(prices)
}
