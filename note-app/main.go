package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"duypn4.dev/note/note"
	"duypn4.dev/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getUserData()
	text := getUserInput("Todo text: ")

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		return
	}

	outputData(todo)
}

func getUserData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(promt string) string {
	fmt.Printf("%v", promt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("failed to save file")
		return err
	}

	fmt.Println("saved file")
	return nil
}
