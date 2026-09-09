package main

import (
	"fmt"
	"test/todo"
)

func main() {

	fmt.Println("hello world")

	todo1 := todo.NewTodo("name", "abcded lorem")

	fmt.Println("todo1:::", todo1)

	nums := []int{1, 2, 3, 4}

	doubleNums := func(nums *[]int) []int {
		result := make([]int, len(*nums))
		for i, v := range *nums {
			result[i] = v * 2
		}
		return result
	}
	fmt.Println(doubleNums(&nums))
	fmt.Println(doubleNums(&nums))
	fmt.Println(nums)

	// scanln data
	var title2, name2 string

	fmt.Println("Enter title 2::: ")
	fmt.Scanln(&title2)

	fmt.Println("Enter name 2:::")
	fmt.Scanln(&name2)

	// create Todo2
	title2, err1 := create("title")(title2)
	name2, err2 := create("name")(name2)
	if err1 != nil || err2 != nil {
		fmt.Println("Error reading file:", err1, err2)
	}
	todo2 := todo.NewTodo(
		title2,
		name2,
	)
	fmt.Printf("todo2:: %v\n", todo2)
}
func create(your_type string) funcType {
	if your_type == "title" {
		return createtitle
	} else {
		return createName
	}
}

type funcType func(string) (string, error)

func createtitle(title string) (string, error) {
	if title == "" {
		return "", fmt.Errorf("Title not found")
	}
	return title, nil
}
func createName(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("Name not found")
	}
	return name, nil
}
