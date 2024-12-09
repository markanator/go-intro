package main

import (
	"bufio"
	"example.com/markanator/user-input/note"
	"example.com/markanator/user-input/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todoObj, err := todo.New(todoText)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
	err = outputData(todoObj)
	if err != nil {
		return
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Todo saved successfully")
	return nil
}

func getNoteData() (string, string) {
	noteTitle := getUserInput("Title:")
	noteContent := getUserInput("Content:")

	return noteTitle, noteContent
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value // clean text to be returned
}
