package functionsarevalue

import "fmt"


func main (){
	fmt.Println("hello world!!")
	numbers := []int{1,2,4,5}
	doubledNumbers := doubleNumber(&numbers)
	fmt.Println(doubledNumbers)
	fmt.Println(numbers)

	/// 1. basic callback function
	doubledNumbersWithCallback := callbackFunc(&numbers, doubleFunc)
	fmt.Println("Print number::::",doubledNumbersWithCallback)

	//2. return function as value from func name 
	moreNums := []int{4,1,2}
	fmt.Println("moreNums::::",moreNums)
	moreNumsProcessed := callbackFunc(&moreNums, handlerManyFunc(&moreNums))
	fmt.Println("moreNumsProcessed::",moreNumsProcessed)

	//3. anonymous function
	
}

func doubleNumber( number *[]int) []int {

	result := make([]int, len(*number))

	for i, v := range *number {
		result[i] = v * v
	}
	return result
}
func tripleNumber (number *[]int) []int {
	result := make([]int, len(*number))
	for i,v := range *number{
		result[i] = v * v * v
	}
	return result
}



// handler callback func type to func 
type funcType func(int) int 

// ly do khai bao *[]int de passing by pointer, de tranh viec copy slice khi truyen vao func
func callbackFunc(number *[]int, callbackfunc funcType) []int{
	 
	result := make([]int, len(*number))

	for i,v := range *number{
		result[i] = callbackfunc(v)
	}

	return result
}
func doubleFunc(v int) int {
	return v * v
}

func tripleFunc(v int) int {
	return v * v * v
}

func handlerManyFunc(nums *[]int ) funcType {
	if (*nums)[0] == 1 {
		return doubleFunc
	} else {
		return tripleFunc
	} 

} 