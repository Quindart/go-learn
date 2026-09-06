package main

import "fmt"


func main (){
	fmt.Println("hello world!!")
	//3. anonymous function
	number := []int{1,2,3}
	transformed := transformedFunc(&number, func(i int) int {
		return i * 2
	})
	fmt.Println("transformed::::",transformed)
}

type funcType func(int) int 
func transformedFunc( nums *[]int, funcTry funcType) []int {
	result := make([]int, len(*nums))
	for i,v := range *nums {
		result[i] = funcTry(v)
	}
	return result
}