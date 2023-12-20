package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang</h1>"))
}

// go mod tidy - removes unwanted packages, refersh go.mod
// go mod verify
// go list
// go build
// go list -m -versions github.com/gorilla/mux  - gives possible versions of the package
// go mod edit -go 1.16 - changes go version in go.mod
// go mod vendor - This command is particularly useful when you want to bundle your project's dependencies along with your source code.
