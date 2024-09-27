package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Content: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("\nContent: %v\n\n", todo.Content)
}

func (todo Todo) Save() error {
	filename := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0644)

}
