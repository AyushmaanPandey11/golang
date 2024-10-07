package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/entity"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	result := add(5, 3)
	fmt.Println("Result is : ", result)
	resultFloat := add(5.4, 3.7)
	fmt.Println("Result is : ", resultFloat)
	resultStr := add("Hello ", "JI")
	fmt.Println("Result is : ", resultStr)

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = SaveAndOutputData(todo)
	if err != nil {
		return
	}
	userNote, err := entity.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = SaveAndOutputData(userNote)
	if err != nil {
		return
	}
}

func add[T string | int | float32 | float64](a, b T) T {
	return a + b
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}
	floatVal, ok := value.(float32)
	if ok {
		fmt.Println("Float : ", floatVal)
		return
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}
}

func SaveAndOutputData(data outputtable) error {
	data.Display()
	return data.Save()
}

func SaveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving the Note Succeded")
	return nil
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
