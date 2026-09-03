package goifelse

import "fmt"

func main(){

	fmt.Println("Welcome back!")
	fmt.Println("1. Check cost")
	fmt.Println("2. Update cost with discount")
	fmt.Println("3. Exit")

	var choice int 
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	
	switch choice {
	case 1:
		fmt.Println("You chose to check cost.")
	case 2:
		fmt.Println("You chose to update cost with discount.")
	case 3:
		fmt.Println("Exiting the program. Goodbye!")
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}