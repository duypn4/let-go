package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"duypn4.dev/note/note"
)

func main() {
	title, content := getUserData()

	note, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	note.Display()

	err = note.Save()
	if err != nil {
		fmt.Println("failed to save file")
	}
	fmt.Println("saved file")
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
