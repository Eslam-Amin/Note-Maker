package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}


func main (){
	title, content := getNoteData()
	userNote, err :=note.New(title, content)
	todoText := getUserInput("Todo Text:")
	if err != nil {
		fmt.Println(err)
		return
	}
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.DisplayTodo()
	
	err = saveData(todo)

	if err != nil {
		return
	}
	fmt.Println("Saving the todo succeeded!")
	userNote.DisplayNote()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data saver) error{
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err 
	}
	fmt.Println("Saving the note succeeded!")

	data.Save()
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")
	return title, content
}


func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)
	reader := bufio.NewReader(os.Stdin)
	// '\n' is the newline character where we press enter, and should be read
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

