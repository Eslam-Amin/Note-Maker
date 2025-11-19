package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)


func main (){
	title, content := getNoteData()
	userNote, err :=note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.DisplayNote()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}
	fmt.Println("Saving the note succeeded!")
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

