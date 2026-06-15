package main

import (
	"fmt"
	"go-struct/user"
)

func main() {
	var quang user.User
	quang.Name = "Quang"
	quang.Age = 20
	quang.Clazz = "Go"

	fmt.Println(quang.Name)
	fmt.Println(quang.Age)
	fmt.Println(quang.Clazz)
	fmt.Printf("Địa chỉ của quang:       %p\n", &quang)
	fmt.Printf("Địa chỉ của quang.Name:  %p\n", &quang.Name)  // Sẽ giống hệt địa chỉ quang
	fmt.Printf("Địa chỉ của quang.Age:   %p\n", &quang.Age)   // Dịch đi một đoạn
	fmt.Printf("Địa chỉ của quang.Clazz: %p\n", &quang.Clazz) // Dịch đi tiếp một đoạn

	var minh = user.User{
		Name:  "Minh",
		Age:   22,
		Clazz: "Go",
	}

	fmt.Println(minh.Name)
	fmt.Println(minh.Age)
	fmt.Println(minh.Clazz)
	fmt.Printf("Địa chỉ của minh:       %p\n", &minh)
	fmt.Printf("Địa chỉ của minh.Name:  %p\n", &minh.Name)  // Sẽ giống hệt địa chỉ minh
	fmt.Printf("Địa chỉ của minh.Age:   %p\n", &minh.Age)   // Dịch đi một đoạn
	fmt.Printf("Địa chỉ của minh.Clazz: %p\n", &minh.Clazz) // Dịch đi tiếp một đoạn
	user.OutputUserDetail(quang)

	userCpy := user.CopyUserStruct(&quang)
	fmt.Println("userCpy", userCpy)
	minh.PrintUser()

	quang.PrintUser()
	quang.UpdateAgeValue(30)

	fmt.Println("Sau khi gọi UpdateAgeValue:")
	quang.PrintUser() // Tuổi vẫn là 20, không đổi

	quang.UpdateAgePointer(30)
	fmt.Println("Sau khi gọi UpdateAgePointer:")
	quang.PrintUser() // Tuổi đã đổi thành 30

	var long *user.User
	long, err := user.New("Long", -25, "Go")
	fmt.Print("long pointer: ", &long)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(long.Name)
	fmt.Println(long.Age)
	fmt.Println(long.Clazz)

}
 