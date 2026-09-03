package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var x int = 10
	fmt.Printf("abcd: %d\n", x)

	// gán biến
	var a, b = 1, 2
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = swap(a, b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
	//
	swap2(a, b)

	// gán biến ẩn
	var m1, _ = GetMoney() // _ -> m2 không phân bổ trong bộ nhớ
	fmt.Printf("m1: %d\n", m1)

	// pham vi bien
	// bieens cuc bo: khai bao trong than ham goi
} 

// 2 cách swap golang
func swap(a, b int) (int, int) {
	var temp = a
	a = b 
	b = temp 
	return a,b
}

func swap2(a,b int ) {
	 a,b = b,a 

	 fmt.Println(a,b)
}

func GetMoney() (int, int) {
	return 1000, 2000
}