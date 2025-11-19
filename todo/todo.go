package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)


type Todo struct{
	Text 	string	`json:"text"`
	CreatedAt time.Time `json:"created_at"`
}


func (todo Todo) Display(){
	fmt.Println(todo)
}

func (todo Todo) Save()error{
	fileName := "todo.json"
	jsonContent, err := json.Marshal(todo)
	
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonContent, 0644)
	
}


func New (content string) (*Todo, error){
	if content == "" {
		return nil, errors.New("title and content are required")
	}

	return &Todo{
		Text: content,
		CreatedAt: time.Now(),
	}, nil 
}