package main

import "fmt"

func main() {
	//? using make function to create a slice

	userNames := make([]string, 2)

	userNames = append(userNames, "ali")
	userNames = append(userNames, "veli")
	userNames = append(userNames, "ahmet")
	userNames = append(userNames, "mehmet")
	fmt.Println(userNames)

	//? using make function to create a map
	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 5
	courseRatings["react"] = 4.5
	courseRatings["flutter"] = 4.7
	fmt.Println(courseRatings)

}
