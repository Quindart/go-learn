package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	ID      int `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func New(id int, title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Title and content cannot be empty")
	}
	fmt.Println("Saving note succeeded !")
	return &Note{
		ID:      id,
		Title:   title,
		Content: content,
	}, nil
}

func (n *Note) Update(title string, content string) error {
	if title == "" || content == "" {
		return errors.New("Title and content cannot be empty")
	}
	n.Title = title
	n.Content = content
	fmt.Println("Note updated successfully !")
	return nil
}

func (n Note) Save() {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshaling note:", err)
		return
	}
	os.WriteFile(fileName, json, 0644)
}
