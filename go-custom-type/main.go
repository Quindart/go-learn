package main

import "fmt"

// custom type definition
type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "John Doe"
	name.log()
}
