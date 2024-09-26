package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/entity"
)

func main() {
	title, content := getNoteData()
	userNote, err := entity.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
	}
	fmt.Println("Saving the Note Succeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter Note Title: ")
	content := getUserInput("Enter Note Content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}
