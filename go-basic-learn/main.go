package main

import (
	"fmt"
	"go-basic-learn/Book"
)

func main() {
	fmt.Println("Hello, World!")

	var myBook = Book.New("The Go Programming Language", "Alan A. A. Donovan", 380)
	fmt.Printf("Book Name: %s\n", myBook.GetName())
	fmt.Printf("Author: %s\n", myBook.GetAuthor())
	fmt.Printf("Pages: %d\n", myBook.GetPages())
}
