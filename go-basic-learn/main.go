package main

import (
	"fmt"
	"go-basic-learn/author"
	"go-basic-learn/book"
	"go-basic-learn/speaker"
)

func main() {
	fmt.Println("Hello, World!")
	var QAuthor = author.New("Quin le")
	var myBook = book.New("The Go Programming Language", QAuthor.GetName(), 30)
	fmt.Printf("Book Name: %s\n", myBook.GetName())
	fmt.Printf("Author: %s\n", myBook.GetAuthor())
	fmt.Printf("Pages: %d\n", myBook.GetPages())
	fmt.Println(QAuthor.Speak())
	var s speaker.ISpeaker = author.New("La le")
	fmt.Println(s.Speak())
}
