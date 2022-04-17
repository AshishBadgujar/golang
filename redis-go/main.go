package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisClient = rdb

	rdb.Ping(ctx)

	http.HandleFunc("/", handleGetVideos)
	http.HandleFunc("/update", handleUpdateVideos)
	http.ListenAndServe(":8080", nil)
}

func handleGetVideos(w http.ResponseWriter, r *http.Request) {
	id, ok := r.URL.Query()["id"]

	if ok {
		videoID := id[0]
		video := getVideo(videoID)

		if video.Id == "" {
			//no videos found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{}"))
			return
		}
		videoBytes, err := json.Marshal(video)
		if err != nil {
			panic(err)
		}
		w.Write(videoBytes)
		return
	}

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
		_, ok := r.URL.Query()["id"]
		if ok {
			var video video
			err = json.Unmarshal(body, &video)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Bad request")
			}
			saveVideo(video)
			return
		}
		var videos []video
		err = json.Unmarshal(body, &videos)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "bad request")
		}
		saveVideos(videos)
		return
	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not supported")
	}
}
