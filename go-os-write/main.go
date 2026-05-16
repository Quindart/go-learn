package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	//0644: Là quyền truy cập cho chủ sở hữu (đọc và ghi)
	//nhóm (đọc) và người khác (đọc)
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {

	var a, b float64
	fmt.Print("Enter a: ")
	fmt.Scanln(&a)
	fmt.Print("Enter b: ")
	fmt.Scanln(&b)

	balance := a + b
	fmt.Printf("Your balance is: %.2f\n", balance)
	writeBalanceToFile(balance)
}
