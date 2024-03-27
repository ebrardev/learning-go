package main

import (
	"fmt"

	"ebrarcode.dev/gonote/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string) {
	title := getUserInput("Note title")

	content := getUserInput("note content")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string

	fmt.Scanln(&value)

	return value
}
