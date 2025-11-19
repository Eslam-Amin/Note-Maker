package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct{
	title			string
	content 	string
	createdAt time.Time
}

func (note *Note) DisplayNote(){
	fmt.Printf("Your note titled: %v has the following content:\n\n %v \n\n", note.title, note.content)
}


func New (title, content string) (*Note, error){
	if title == "" || content == "" {
		return nil, errors.New("title and content are required")
	}

	return &Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil 
}