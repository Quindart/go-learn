package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


const balanceFileName = "../go-os-write/balance.txt"

func getBalanceFromFile() error {

	data, _ := os.ReadFile(balanceFileName)
	balanceText := string(data)
	balanceFloat64, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return errors.New("Failed to parse balance from file: " + err.Error())
	}

	fmt.Printf("Your balance from file is: %f\n", balanceFloat64)
	fmt.Printf("Your balance from file is: %s\n", balanceText)
	return nil
}

func main() {
	getBalanceFromFile()
}
