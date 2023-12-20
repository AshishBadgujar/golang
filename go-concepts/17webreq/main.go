package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to golang server...")
	// PerformGetRequest("http://localhost:8000/get")
	// PerformPostRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerformGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
func PerformPostRequest(url string) {
	requestBody := strings.NewReader(`
	{
		"coursename":"let's go with golang",
		"price":0,
		"platform":"youtube"
	}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest(myurl string) {
	// formdata

	data := url.Values{}

	data.Add("firstname", "Ashish")
	data.Add("lastname", "Badgujar")
	data.Add("email", "ashish@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
