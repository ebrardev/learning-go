package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ebrarcode.dev/interface/todo"
)

func main() {

	todoText := getUserInput("Todo content:")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("error saving todo: ", err)
		return
	}
	fmt.Println("todo saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
