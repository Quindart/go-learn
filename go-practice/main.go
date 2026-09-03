package main

import "fmt"

func main() {

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
	hobbies := [3]string{"read","write","speak"}
	fmt.Println("Hobbies: ", hobbies)
// 2) Also output more data about that array:
//		- The first element (standalone)
	fmt.Println("First hobby: ", hobbies[0])
//		- The second and third element combined as a new list
	fmt.Println("second, third", hobbies[1:3]) // 1,2
// 3) Create a slice based on the first element that contains
//		the first and second elements.
	fmt.Println("first, second", hobbies[0:2]) // 0,1
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
	slice := hobbies[0:2]
	slice = hobbies[1:3]
	fmt.Println("Re-sliced: ", slice)
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courses := []string{"Learn Go", "Build a project"}
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courses[1] = "Build xcode"
	courses = append(courses, "Build a Microservice")
	fmt.Println(courses)
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
	type Product struct {
		Title string
		ID    int
		Price float64
	}
	products := []Product{
		{Title: "Product 1", ID: 1, Price: 19.99},
		{Title: "Product 2", ID: 2, Price: 29.99},
	}
	products = append(products, Product{Title: "Product 3", ID: 3, Price: 39.99})
	fmt.Println(products)
}
