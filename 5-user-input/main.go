package main

import (
	"bufio"
	"example.com/markanator/user-input/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Note saved successfully")
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
