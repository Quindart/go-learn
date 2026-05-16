package fileos

import 
(
	"fmt"
	"os"
)

func WriteToFile(result1, result2, result3 float64) {
	balanceText := fmt.Sprint("EBT: ", result1, "\n", "Profit: ", result2, "\n", "Profit Margin: ", result3, "%")
	error := os.WriteFile("result.txt", []byte(balanceText), 0644) // 00644: read/write for owner, read for group and others

	if error != nil {
		fmt.Println("Error writing to file:", error)
	}
}

func ReadFromFile() {
	content, err := os.ReadFile("result.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:\n", string(content))
}
