package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var Notes []Note
const NOTES_FILE = "notes.json"

type Note struct{
	Title			string `json:"title"`
	Content 	string	`json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note *Note) Display(){
	fmt.Printf("Your note titled: %v has the following content:\n\n %v \n\n", note.Title, note.Content)
}

func (note Note) Save() error{
	fmt.Println("NOTES #0: ",Notes)
	listOfNotes, err := loadNotes()
	if err != nil {
		return err
	}
	
	listOfNotes = append(listOfNotes, note)
	fmt.Println("list of notes: ",listOfNotes)
	Notes = listOfNotes
	fmt.Println("NOTES #1: ",Notes)
	jsonContent, err := json.Marshal(listOfNotes)
	if err != nil {
		return err
	}
	return os.WriteFile(NOTES_FILE, jsonContent, 0644)
}

func loadNotes ()([]Note, error){
	var list []Note
	if len(Notes) > 0 {
		return Notes, nil
	}
	data, err := os.ReadFile(NOTES_FILE)
	if err != nil {
		return nil, err
	}
	if len(data) ==0 {
		Notes = list
		return list, nil
	}
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}

	Notes = list
	return list, nil
}

func New (title, content string) (*Note, error){
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}

	return &Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil 
}