package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api/", handleGetTodos)
	http.HandleFunc("/api/submit", handleSaveTodo)
	http.HandleFunc("/api/delete", handleDeleteTodo)
	fmt.Println("Starting server on the port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleGetTodos(w http.ResponseWriter, r *http.Request) {
	todos := getTodos()
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	w.Write(todoBytes)
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		todos := getTodos()
		newTodos := todos[:0]
		for _, item := range todos {
			if item.Id != id {
				newTodos = append(newTodos, item)
			}
		}
		saveTodo(newTodos)
		// w.WriteHeader(200)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not supported")
	}

}

func handleSaveTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var myTodo todo
		err = json.Unmarshal(body, &myTodo)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "bad request")
		} else {
			fmt.Println(myTodo)
			todos := getTodos()
			todos = append(todos, myTodo)
			saveTodo(todos)
		}
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not supported")
	}

}
