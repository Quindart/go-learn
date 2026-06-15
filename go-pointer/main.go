// Panic: runtime error: invalid memory address or nil pointer dereference.
// This error occurs when you try to access a memory location that has not been allocated or is nil. In the above code, we are trying to dereference a nil pointer, which causes the panic. To fix this issue, we need to ensure that the pointer is properly initialized before dereferencing it.
// Nếu tờ giấy ghi "Số 123 đường ABC", nghĩa là con trỏ đang trỏ tới một ô nhớ cụ thể. Bạn có thể đến đó để lấy dữ liệu.
// Nếu tờ giấy để trống hoàn toàn, đó chính là nil. Tờ giấy tồn tại, nhưng nó không dẫn bạn đến bất kỳ ngôi nhà (ô nhớ) nào cả.
// For example, the null value of an int variable is 0. Of a float64, it would be 0.0. Of a string, it's "".
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello go")
	var a int = 10
	var b *int = &a
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a)
	fmt.Printf("Value of b: %d\n", *b)
	fmt.Printf("Address of b: %p\n", b)
	// a = 40
	// fmt.Printf("Value of a: %d\n", a)
	// fmt.Printf("Address of a: %p\n", &a)
	// fmt.Printf("Value of b: %d\n", *b)
	// fmt.Printf("Address of b: %p\n", b)
	// *b = 20
	// fmt.Printf("Value of a: %d\n", a)
	// fmt.Printf("Address of a: %p\n", &a)
	// fmt.Printf("Value of b: %d\n", *b)
	// fmt.Printf("Address of b: %p\n", b)

	var age = 30 // passing by value
	var agePoint = &age // passing by reference (pointer)

	fmt.Printf("Value of age: %d\n", age)
	fmt.Printf("Address of age: %p\n", &age)
	fmt.Printf("Value of agePoint: %d\n", *agePoint) // passing the value of age through the pointer
	fmt.Printf("Address of agePoint: %p\n", agePoint) // passing the address of age through the pointer	
	fmt.Printf("Adult years: %d\n", getAdultYears(agePoint)) // passing the pointer to the function, which will dereference it to get the value of age
}

func getAdultYears(agePoint *int) int {
	var ageVal = *agePoint
	if ageVal >= 18 {
		return ageVal - 18
	}
	return 0
}
 