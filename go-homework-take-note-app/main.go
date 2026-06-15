package main

import (
	"errors"
	"fmt"
	"go-homework-take-note-app/note"
	"strconv"
)

func main() {

	fmt.Println("Welcome to note app ")

	var note1 *note.Note

	note1, err := note.New(1, "First Note", "This is the content of the first note.")
	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}
	fmt.Print("Note created:", &note1)

	fmt.Print("Saved file!")
	note1.Save()

	var choice int
	fmt.Println("Please select an option:")
	fmt.Println("1. Create a new note")
	fmt.Println("2. View all notes")
	fmt.Println("3. Update a note")
	fmt.Println("4. Delete a note")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	viewMenu(choice)

	//
}

func viewMenu(choice int) {
	switch choice {
	case 1:
		fmt.Println("Create a new note")
		noteIDStr, err := enterInput("Enter note ID: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		noteID, err := strconv.Atoi(noteIDStr)
		if err != nil {
			fmt.Println("Error: note ID must be a number")
			break
		}
		noteTitle, err := enterInput("Enter note title: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		noteContent, err := enterInput("Enter note content: ")
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		n, err := note.New(noteID, noteTitle, noteContent)
		if err != nil {
			fmt.Println("Error creating note:", err)
			break
		}
		fmt.Println("Note created:", n)
	case 2:
		fmt.Println("View all notes")
	case 3:
		fmt.Println("Update a note")
	case 4:
		fmt.Println("Delete a note")
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
