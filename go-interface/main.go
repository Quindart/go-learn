package main

import (
	"errors"
	"fmt"
	"go-interface/Todo"
	"strconv"
)



func main() {

	fmt.Println("Welcome to Todo app ")

	var todo1 *Todo.Todo

	todo1, err := Todo.New(1, "First Todo", "This is the content of the first Todo.")
	if err != nil {
		fmt.Println("Error creating Todo:", err)
		return
	}
	fmt.Print("Todo created:", &todo1)

	fmt.Print("Saved file!")
	todo1.Save()

	var choice int
	fmt.Println("Please select an option:")
	fmt.Println("1. Create a new Todo")
	fmt.Println("2. View all Todos")
	fmt.Println("3. Update a Todo")
	fmt.Println("4. Delete a Todo")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	viewMenu(choice)
}

func viewMenu(choice int) {
	switch choice {
	case 1:
		fmt.Println("Create a new Todo")
		TodoIDStr, err := enterInput("Enter Todo ID: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		TodoID, err := strconv.Atoi(TodoIDStr)
		if err != nil {
			fmt.Println("Error: Todo ID must be a number")
			break
		}
		TodoTitle, err := enterInput("Enter Todo title: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		TodoContent, err := enterInput("Enter Todo content: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		n, err := Todo.New(TodoID, TodoTitle, TodoContent)
		if err != nil {
			fmt.Println("Error creating Todo:", err)
			break
		}
		fmt.Println("Todo created:", n)
	case 2:
		fmt.Println("View all Todos")
	case 3:
		fmt.Println("Update a Todo")
	case 4:
		fmt.Println("Delete a Todo")
	default:
		fmt.Println("Invalid option, please try again.")
	}
}

func enterInput(label string) (string, error) {

	var text string
	fmt.Print(label)
	fmt.Scan(&text)

	if text == "" {
		return "", errors.New("Input cannot be empty")
	}
	return text, nil
}
