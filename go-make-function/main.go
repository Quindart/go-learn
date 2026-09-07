package main

import "fmt"

// alias: you can create an alias for a type using the type keyword.
//  This can be useful for creating more descriptive names for types, or for creating new types that are based on existing types.
type floatMap map[string]float64

func (f floatMap) output(data floatMap) {
	fmt.Println("mapp data::::", data)
}
func main() {

	username := make([]string, 2, 5) // tạo một vùng chứa string với 2 phần tử nhưng chứa tối đa 5 phần tử

	username[0] = "John"
	username[1] = "Joe"

	// username[2]= "Doe" // lỗi vì chỉ có 2 phần tử được khởi tạo, nhưng có thể chứa tối đa 5 phần tử
	// username[3]= "Smith" // lỗi vì chỉ có 2 phần tử được khởi tạo, nhưng có thể chứa tối đa 5 phần tử

	username = append(username, "Alice")   // thêm phần tử vào slice, nếu vượt quá capacity thì sẽ tự động tăng kích thước slice
	username = append(username, "Bob")     // thêm phần tử vào slice, nếu vượt quá capacity thì sẽ tự động tăng kích thước slice
	username = append(username, "Charlie") // thêm phần tử vào slice, nếu vượt quá capacity thì sẽ tự động tăng kích thước slice
	username = append(username, "David")   // thêm phần tử vào slice, nếu vượt quá capacity thì sẽ tự động tăng kích thước slice

	fmt.Println("username::::", username)

	demoMap := make(floatMap, 2)
	demoMap["a"] = 1.0
	demoMap["b"] = 2.0

	demoMap.output(demoMap)

	var s []int // s là nil

	fmt.Println(s == nil) // true
	fmt.Println(len(s))   // Output: 0

	s = append(s, 10)     // Hoạt động tốt! Go tự động cấp phát bộ nhớ
	fmt.Println(s)

	// loop in array
	for i, value := range username {
		fmt.Println(value)
		fmt.Println("index::::", i)
	}

}
