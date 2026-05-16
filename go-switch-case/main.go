package main
import "fmt"

func main(){

	var choice int
	fmt.Print("Enter your choice (1/0): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("You chose true.")
	case 0:
		fmt.Println("You chose false.")
	default:
		fmt.Println("Invalid choice. Please enter 1 for true or 0 for false.")
	}
}