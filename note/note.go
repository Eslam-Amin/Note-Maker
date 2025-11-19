package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)


type Note struct{
	Title			string
	Content 	string
	CreatedAt time.Time
}

func (note *Note) DisplayNote(){
	fmt.Printf("Your note titled: %v has the following content:\n\n %v \n\n", note.Title, note.Content)
}

func (note Note) Save()error{
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	jsonContent, err := json.Marshal(note)
	
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
	
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