package main

import "fmt"

func main() {

	value := fibonacci(5)
	fmt.Println("value::::", value)

	// numbers list
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(numbers)
	fmt.Println("sum::::", sum)

	fmt.Println("summ variadic::::", variadicfunc(1, 2, 3, 4, 5))

	// merge array
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	merge := append(a, b...) // this is a merged array
	fmt.Println("merge::::", merge)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) * n
}

// sumup
func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// variadic function
func variadicfunc(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

