package main

import "fmt"

func main() {
	fmt.Println("hello world!!")
	//3. anonymous function
	number := []int{1, 2, 3}
	transformed := transformedFunc(&number, func(i int) int {
		return i * 2
	})
	fmt.Println("transformed::::", transformed)

	// khi gán double trùng tên với 1 function khác thì double đó ưu tiên hàm được gán dưới dạng annonymos function trước
	double := double3(4)
	fmt.Println("transformedByClosure::::", double)
}

type funcType func(int) int

func transformedFunc(nums *[]int, funcTry funcType) []int {
	result := make([]int, len(*nums))
	for i, v := range *nums {
		result[i] = funcTry(v)
	}
	return result
}

func double(v int) int {
	return v * v
}

func double3(v int) int {
	return v * v * 3
}

// todo tranformfunction closure. 
//? Closure (hàm bao đóng) trong Golang là một hàm ẩn danh có tham chiếu đến các biến nằm bên ngoài phạm vi của nó.
// [1] (https://translate.google.com/translate?u=https://go.dev/tour/moretypes/25&hl=vi&sl=en&tl=vi&client=sge),
// [2] (https://translate.google.com/translate?u=https://gobyexample.com/closures&hl=vi&sl=en&tl=vi&client=sge)
func createTransform(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
