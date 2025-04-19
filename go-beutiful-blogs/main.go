package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api/", handleGetContacts)
	http.HandleFunc("/api/submit", handleSaveContacts)
	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleGetContacts(w http.ResponseWriter, r *http.Request) {
	contacts := getContacts()
	contactByets, err := json.Marshal(contacts)
	if err != nil {
		panic(err)
	}
	w.Write(contactByets)
}

func handleSaveContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var myContact contact
		err = json.Unmarshal(body, &myContact)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "bad request")
		} else {
			fmt.Println(myContact)
			contacts := getContacts()
			contacts = append(contacts, myContact)
			saveContact(contacts)
		}
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not supported")
	}

}
