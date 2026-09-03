package main

import "fmt"

func main() {
	arr := [40]int{1, 2, 3, 4, 5} // 8 bytes for each int, total 40 bytes for the array - mallocated memory for 40 integers, but only 5 are initialized

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Int value: %v\n", arr[i])
		fmt.Printf("Int pointer: %v\n", &arr[i])
	}
	fmt.Printf("Length of array: %v\n", len(arr))

	// without strings
	//Note*: Malloc (viết tắt của Memory Allocation) là một hàm trong ngôn ngữ lập trình C/C++.
	// Chức năng chính của nó là cấp phát một vùng nhớ động trên vùng nhớ Heap với kích thước (tính bằng byte) được chỉ định trước
	var stringNames = [2]string{"John", "Doe"}
	var stringLocation = &stringNames
	fmt.Printf("String name: %v\n", stringNames)
	fmt.Printf("String pointer: %v\n", *stringLocation)


	fmt.Printf("[1:3]: %v\n", arr[1:3]) // -> 2, 3
	fmt.Printf("[:3]: %v\n", arr[:3])   // -> 1, 2, 3
	fmt.Printf("[:3]: %v\n", arr[:3])   // -> 1, 2, 3

	//Note*: [1:3] => slice by value, not by reference
	prices := [4]float64{19.99, 29.99, 39.99, 49.99}
	fmt.Println("Prices:", prices)

	for i := 0; i < len(prices); i++ {
		fmt.Printf("Price: %v : %v\n", &prices[i], prices[i])
	}

	feat_price := prices[1:]
	high_price := feat_price[:1]
	for i := 0; i < len(high_price); i++ {
		fmt.Printf("High price: %v : %v\n", &high_price[i], high_price[i])
	}
	for i := 0; i < len(feat_price); i++ {
		fmt.Printf("Featured price: %v : %v\n", &feat_price[i], feat_price[i])
	}
	// unpacking list in golang = destructuring list in javascript
	list123 := []int{1,2,3}
	list456 := []int{4,5,6}
	merge := append(list123, list456...)
	fmt.Println(merge)
	
}
