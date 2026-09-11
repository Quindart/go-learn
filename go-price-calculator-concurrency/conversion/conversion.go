package conversion

import (
	"fmt"
	"strconv"
)

func StringToFloat(lines []string) ([]float64, error) {
	result := make([]float64, len(lines))
	for index, line := range lines {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Failed to parse string to float64 err:::", err)
			return nil, err
		}
		result[index] = value
	}
	return result, nil
}
