package main

import "fmt"

func main() {
	var accountBalance = 1000.0
	for i := 1; i < 5; i++ {
		fmt.Printf("Welcome to go bank!")
		fmt.Printf("Your current balance is: %.2f\n", accountBalance)
		var depositAmount float64
		fmt.Print("Enter the amount to deposit: ")
		fmt.Scanln(&depositAmount)
		accountBalance += depositAmount
	}

	for accountBalance < 2000 {
		fmt.Printf("Your current balance is: %.2f\n", accountBalance)
		var depositAmount float64
		fmt.Print("Enter the amount to deposit: ")
		fmt.Scanln(&depositAmount)
		accountBalance += depositAmount
	}
	fmt.Printf("Your final balance is: %.2f\n", accountBalance)

}
