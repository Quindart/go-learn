package main

import "fmt"

func main() {

	a := getUserInput("Enter a: ")
	b := getUserInput("Enter b: ")

	total, pro := handleCaculate(a,b)
	fmt.Printf("Total result: %v, Product result: %v\n", total, pro)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scanln(&userInput)
	return userInput
}

func handleCaculate(a,b float64) (float64, float64) {
	return a + b, a * b
}
