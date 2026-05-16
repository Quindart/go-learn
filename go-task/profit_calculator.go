package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1. Validate user input
// 2. store calculated results into file

func main() {

	revenue, err := getInput("Revenue: ")
	expenses, err := getInput("Expenses: ")
	tax, err := getInput("Tax: ")
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ebt, profit, profitMargin := calculateFinancials(revenue, expenses, tax)

	fmt.Printf("Revenue: %.2f\nExpenses: %.2f\nTax: %.2f\nEBT: %.2f\nProfit: %.2f\nProfit Margin: %.2f%%\n", revenue, expenses, tax, ebt, profit, profitMargin)
	fmt.Println("Results have been written to result.txt")
	writeToFile(ebt, profit, profitMargin)

	fmt.Println("Reading from file...")
	readFromFile()
}

func calculateFinancials(revenue, expenses, tax float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt - tax
	profitMargin := (profit / revenue) * 100

	return ebt, profit, profitMargin
}

func getInput(text string) (float64, error) {

	var textFloat float64
	fmt.Print(text)
	fmt.Scan(&textFloat)

	if textFloat < 0 {
		return 0, errors.New("Value cannot be negative")
	}

	return textFloat, nil
}

func writeToFile(result1, result2, result3 float64) {
	balanceText := fmt.Sprint("EBT: ", result1, "\n", "Profit: ", result2, "\n", "Profit Margin: ", result3, "%")
	error := os.WriteFile("result.txt", []byte(balanceText), 0644) // 00644: read/write for owner, read for group and others

	if error != nil {
		fmt.Println("Error writing to file:", error)
	}
}

func readFromFile() {
	content, err := os.ReadFile("result.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(content))
}
