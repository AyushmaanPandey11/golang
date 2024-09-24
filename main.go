package main

import (
	"fmt"

	"example.com/note/entity"
)

func main() {
	title, content := getNoteData()
	userNote, err := entity.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note Title: ")
	content := getUserInput("Enter Note Content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
