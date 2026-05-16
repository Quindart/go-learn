package main

import (
	"fmt"
)

func main(){
	fmt.Println("Hello go")
	var a int = 10
	var b *int = &a
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value of b: %d\n", *b)
	fmt.Printf("Address of b: %p\n", b)
	a = 40
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value of b: %d\n", *b)
	fmt.Printf("Address of b: %p\n", b)
	*b = 20
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value of b: %d\n", *b)
	fmt.Printf("Address of b: %p\n", b)
}