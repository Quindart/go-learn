package main

import "fmt"

// biến toàn cục
var globalVar = 1232131 

func main() {
	// khai bao bien cục bộ
	var x int = 10
	var y = 20
	z := x + y
	globalVar = x + 20
	// in ra gia tri bien
	println("x:", x)
	println("y:", y)
	println("z:", z) 
	println("globalVar:", globalVar)
	printGlobalVar()
}
// go cho phép khai báo biến toàn cục và biến cục bộ cùng tên 
// => thì nó ưu tiên biến cục bộ hơn
// quindart@MacBook-Pro-cua-Quindart variable-scope % go run main.go
// x: 10
// y: 20
// z: 30
// globalVar: 30
// globalVar: 30

 func printGlobalVar() {
	fmt.Println("globalVar:", globalVar)
}
