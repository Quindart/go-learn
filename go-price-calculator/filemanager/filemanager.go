package filemanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path) //"./prices.txt"
	if err != nil {
		fmt.Println("Failed to open file err:::", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	scanner.Scan()
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file err:::", err)
		file.Close()
		return nil, err
	}
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("Failed to encoder", err)
	}
	file.Close()
	return nil
}
