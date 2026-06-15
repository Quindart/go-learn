package main

import "fmt"

func printSomething(s interface{}) { //Empty Interface   interface{} = any
	fmt.Println(s)
}

func onSwitchType(demo_type interface{}) {
	switch v := demo_type.(type) {
	case string:
		fmt.Printf("The type of demo_type is string and the value is: %s\n", v)
	case int:
		fmt.Printf("The type of demo_type is int and the value is: %d\n", v)

	case float32:
		fmt.Printf("The type of demo_type is float32 and the value is: %f\n", v)
	default:
		fmt.Printf("The type of demo_type is unknown and the value is: %v\n", v)
	}
}

func main() {
	printSomething("Hello, World!")
	printSomething(42)
	onSwitchType("Hello, World!")
	onSwitchType(42)
	onSwitchType(3.14)
}
