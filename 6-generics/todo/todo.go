package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	formatted, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, formatted, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: content,
	}, nil
}
