package main

import (
	"encoding/json"
	"io/ioutil"
)

type contact struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func getContacts() (contacts []contact) {
	fileBytes, err := ioutil.ReadFile("./contacts.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &contacts)
	if err != nil {
		panic(err)
	}
	return contacts
}

func saveContact(contacts []contact) {
	contactBytes, err := json.Marshal(contacts)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./contacts.json", contactBytes, 0644)
	if err != nil {
		panic(err)
	}
}
