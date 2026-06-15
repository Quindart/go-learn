package Todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type todoInterface interface {
	addTodo(todo string)
	removeTodo(id int)
	getTodos() []string
}

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(id int, title string, content string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	fmt.Println("Saving Todo succeeded !")
	return &Todo{
		ID:      id,
		Title:   title,
		Content: content,
	}, nil
}

func (n *Todo) Update(text string) error {
	if text == "" {
		return errors.New("Text cannot be empty")
	}
	n.Title = text
	fmt.Println("Todo updated successfully !")
	return nil
}

func (n Todo) Save() {
	var fileName = "todo.json"
	data, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling Todo:", err)
		return
	}
	os.WriteFile(fileName, data, 0644)
}
