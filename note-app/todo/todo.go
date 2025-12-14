package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Todo has the text: %v", todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
