package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"swimming", "running", "reading"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0:1])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	myHobbies := hobbies[0:2]
	fmt.Println(myHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	fmt.Println(cap(myHobbies))

	myHobbies = myHobbies[1:3]
	fmt.Println(myHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)

	courseGoals := []string{"learn go", "learn react"}
	fmt.Println(courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "learn react native"
	courseGoals = append(courseGoals, "learn flutter")
	fmt.Println(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{"product1", 1, 10.99},
		{"product2", 2, 20.99},
		{"product3", 3, 30.99},
	}
	fmt.Println(products)
}
