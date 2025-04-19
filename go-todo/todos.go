package main

import (
	"encoding/json"
	"io/ioutil"
)

type todo struct {
	Id   string `json:"id"`
	Todo string `json:"todo"`
}

func getTodos() (todos []todo) {
	fileBytes, err := ioutil.ReadFile("./todo.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &todos)
	if err != nil {
		panic(err)
	}
	return todos
}

func saveTodo(todos []todo) {
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./todo.json", todoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
