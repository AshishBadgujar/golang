package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleGetVideos)
	http.HandleFunc("/update", handleUpdateVideos)
	http.ListenAndServe(":8080", nil)
}

func handleGetVideos(w http.ResponseWriter, r *http.Request) {
	videos := getVideos()
	videosBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}
	w.Write(videosBytes)
}
func handleUpdateVideos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "bad request")
		}
		saveVideos(videos)
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not supported")
	}
}
